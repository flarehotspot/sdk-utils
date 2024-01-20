(function (window) {
  window.apiv1 = {
    HelpersPath: function (pkg) {
      var url = '{{ .Helpers.UrlForMuxRoute "admin.helperjs" "pkg" "PKG" }}';
      return url.replace('PKG', pkg);
    },
    ApiPath: function (pkg) {
      var url = '{{ .Helpers.UrlForMuxRoute "api.apijs" "pkg" "PKG" }}';
      return url.replace('PKG', pkg);
    }
  };

  function HasAuthToken() {
    var segmnts = document.cookie.split(';');
    var hastoken = false;
    for (var i = 0; i < segmnts.length; i++) {
      var seg = segmnts[i].split('=');
      if (seg[0].trim() === 'auth-token' && seg[1].length > 0) {
        hastoken = true;
        break;
      }
    }
    return hastoken;
  }

  function VueLoader(vueFile) {
    return function (resolve) {
      return require(['vue!' + vueFile], resolve);
    };
  }

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
            requireAuth: true
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
      var hastoken = HasAuthToken();

      if (
        to.matched.some(function (route) {
          return route.meta.requireAuth;
        })
      ) {
        hastoken ? next() : next({ name: 'login' });
      }

      if (
        hastoken &&
        to.matched.some(function (route) {
          return route.meta.requireNoAuth;
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
      } else {
        return next();
      }
    });

    // end routes

    var app = new Vue({
      router: router,
      mounted: function () {
        var self = this;
        // handle unauthorized requests
        window.BasicHttp.onUnauthorized = function () {
          var pending = self.$router.history.pending || {};
          var current = self.$router.history.current;
          if (current.name != 'login' && pending.name != 'login') {
            router.push({ name: 'login' });
          }
        };
      }
    });
    app.$mount('#app');
  });
})(window);
