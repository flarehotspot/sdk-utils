(function (window) {
  window.apiv1 = {
    HelperPath: function (pkg) {
      var helperJsURL = '{{ .Data.HelperJsURL }}';
      return helperJsURL.replace('PKG', pkg);
    }
  };

  function VueLoader(vueFile) {
    return function (resolve) {
      return require(['vue!' + vueFile], resolve);
    };
  }

  require.config({
    paths: {
      vue: '{{ .Helpers.AssetPath "libs/requirejs-vue-1.1.5.min" }}',
      json: '{{ .Helpers.AssetPath "libs/requirejs-json-0.4.0.min" }}',
      image: '{{ .Helpers.AssetPath "libs/requirejs-image-0.2.2.min" }}',
      text: '{{ .Helpers.AssetPath "libs/requirejs-text-2.0.5.min" }}'
    },
    config: {
      vue: {
        css: 'inject',
        templateVar: 'template'
      }
    }
  });

  require(['{{ .Helpers.AssetPath "services/http.js" }}'], function (http) {
    var Vue = window.Vue;
    var VueRouter = window.VueRouter;

    // start configs --------------------------------------------
    var routesJson = JSON.parse('{{ .Data.Routes }}');
    var themeLayoutComponent = '{{ .Data.Theme.LayoutComponent }}';
    var themeIndexComponent = '{{ .Data.Theme.IndexComponent }}';
    var themeLoginComponent = '{{ .Data.Theme.LoginComponent }}';
    // end configs --------------------------------------------

    // start routes
    var routes = [
      {
        path: '/',
        name: 'theme-index',
        component: VueLoader(themeIndexComponent)
      }
    ];

    for (var i = 0; i < routesJson.length; i++) {
      var r = routesJson[i];
      routes.push({
        path: r.path,
        name: r.name,
        component: VueLoader(r.component)
      });
    }

    var router = new VueRouter({
      routes: [
        {
          path: '/',
          name: 'theme-layout',
          component: VueLoader(themeLayoutComponent),
          children: routes
        },
        {
          path: '/login',
          name: 'login',
          component: VueLoader(themeLoginComponent)
        }
      ]
    });

    require.onError = function (err) {
      console.error(err);
    };

    // end routes

    var app = new Vue({ router: router });
    app.$mount('#app');

    // check if user is authenticated, if not redirect to login page
    http.getJson(
      '{{ .Helpers.UrlForMuxRoute "auth.is-authenticated" }}',
      function (err) {
        if (err) {
          if (err.status === 401) {
            router.push({ name: 'login' });
          } else {
            console.error(err);
          }
        } else {
          router.push({ name: 'theme-index' });
        }
      }
    );
  });
})(window);
