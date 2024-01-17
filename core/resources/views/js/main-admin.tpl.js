(function (window) {
  window.apiv1 = {
    HelperPath: function (pkg) {
      var url = '{{ .Helpers.UrlForMuxRoute "admin.helperjs" "pkg" "PKG" }}';
      return url.replace('PKG', pkg);
    },
    ApiPath: function (pkg) {
      var url = '{{ .Helpers.UrlForMuxRoute "api.apijs" "pkg" "PKG" }}';
      return url.replace('PKG', pkg);
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

  require([
    '{{  .Helpers.UrlForMuxRoute "api.apijs" "pkg" "com.flarego.core" }}'
  ], function (api) {
    var Vue = window.Vue;
    var VueRouter = window.VueRouter;

    var routesJson = JSON.parse('{{ .Data.Routes }}');
    var themeLayoutComponent = '{{ .Data.Theme.LayoutComponentPath }}';
    var themeLoginComponent = '{{ .Data.Theme.LoginComponentPath }}';

    // start routes
    var routes = [];
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
          children: routes,
          meta: {
            requiresAuth: true
          }
        },
        {
          path: '/login',
          name: 'login',
          component: VueLoader(themeLoginComponent),
          meta: {
            requireNoAuth: true
          }
        }
      ]
    });

    router.beforeEach(function (to, _, next) {
      if (
        to.matched.some(function (record) {
          return record.meta.requiresAuth;
        })
      ) {
        api.Auth.IsAuthenticated()
          .then(function () {
            next();
          })
          .catch(function (err) {
            console.error(err);
            next({ name: 'login' });
          });
      }

      if (
        to.matched.some(function (record) {
          return record.meta.requireNoAuth;
        })
      ) {
        api.Auth.IsAuthenticated()
          .then(function () {
            next({ name: 'theme-index' });
          })
          .catch(function (err) {
            console.error(err);
            next();
          });
      }
    });

    require.onError = function (err) {
      console.error(err);
    };

    // end routes

    var app = new Vue({ router: router });
    app.$mount('#app');
  });
})(window);
