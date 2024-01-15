(function () {
  require.config({
    baseUrl: '/assets/{{ .Data.AssetsVersion }}',
    paths: {
      vue: '{{ .Data.CoreApi.Pkg }}/libs/requirejs-vue-1.1.5.min',
      json: '{{ .Data.CoreApi.Pkg }}/libs/requirejs-json-0.4.0.min',
      image: '{{ .Data.CoreApi.Pkg }}/libs/requirejs-image-0.2.2.min',
      text: '{{ .Data.CoreApi.Pkg }}/libs/requirejs-text-2.0.5.min'
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
    var portalThemeLayoutComponent = '{{ .Data.PortalTheme.LayoutComponent }}';
    var portalThemeIndexComponent = '{{ .Data.PortalTheme.IndexComponent }}';
    // end configs --------------------------------------------

    function vueLoader(vueFile) {
      if (vueFile.charAt(0) === '/') {
        vueFile = vueFile.substring(1);
      }

      return function (resolve) {
        return require(['vue!' + vueFile], resolve);
      };
    }

    // start routes
    var routes = [
      {
        path: '/',
        name: 'portal-index',
        component: vueLoader(portalThemeIndexComponent)
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
          name: 'portal-theme',
          component: vueLoader(portalThemeLayoutComponent),
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
