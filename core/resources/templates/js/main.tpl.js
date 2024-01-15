(function () {
  require.config({
    paths: {
      vue: '{{ .Helpers.AssetPath "libs/requirejs-vue-1.1.5.min" }}',
      json: '{{ .Helpers.AssetPath "libs/requirejs-json-0.4.0.min" }}',
      image: '{{ .Helpers.AssetPath "libs/requirejs-image-0.2.2.min" }}',
      text: '{{ .Helpers.AssetPath "libs/requirejs-text-2.0.5.min" }}',
    },
    config: {
      vue: {
        css: 'inject',
        templateVar: 'template'
      }
    }
  });

  require([], function () {
    var Vue = window.Vue;
    var VueRouter = window.VueRouter;

    // start configs --------------------------------------------
    var routesJson = JSON.parse('{{ .Data.Routes }}');
    var themeLayoutComponent = '{{ .Data.Theme.LayoutComponent }}';
    var themeIndexComponent = '{{ .Data.Theme.IndexComponent }}';
    // end configs --------------------------------------------

    function vueLoader(vueFile) {
      return function (resolve) {
        return require(['vue!' + vueFile], resolve);
      };
    }

    // start routes
    var routes = [
      {
        path: '/',
        name: 'theme-index',
        component: vueLoader(themeIndexComponent)
      }
    ];

    for (var i = 0; i < routesJson.length; i++) {
      var r = routesJson[i];
      routes.push({
        path: r.path,
        name: r.name,
        component: vueLoader(r.component)
      });
    }

    var router = new VueRouter({
      routes: [
        {
          path: '/',
          name: 'theme-layout',
          component: vueLoader(themeLayoutComponent),
          children: routes
        }
      ]
    });
    // end routes

    var app = new Vue({
      router: router
    });

    app.$mount('#app');
  });
})();
