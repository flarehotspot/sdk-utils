/**
 * @file             : router.tpl.js
 * @author           : Adones Pitogo <adones.pitogo@adopisoft.com>
 * Date              : Jan 19, 2024
 * Last Modified Date: Jan 20, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */

(function ($flare) {
  function VueLazyLoad(vueFile) {
    return function (resolve) {
      return require(['vue!' + vueFile], resolve);
    };
  }

  function transformRoutes(routes) {
    var newRoutes = [];
    for (var i = 0; i < routes.length; i++) {
      var r = routes[i];
      var route = {
        name: r.name,
        path: r.path,
        component: VueLazyLoad(r.component),
        meta: r.meta,
        props: true
      };

      if (r.children) {
        route.children = transformRoutes(r.children);
      }

      newRoutes.push(route);
    }
    return newRoutes;
  }

  var VueRouter = window.VueRouter;
  var routes = JSON.parse('<% .Data.Routes %>');
  console.log(routes);

  routes = transformRoutes(routes);

  console.log(routes);

  var appRouter = new VueRouter({ routes: routes });
  $flare.router = appRouter;

  appRouter.beforeEach(function (to, _, next) {
    var hastoken = $flare.auth.hasAuthToken();
    console.log('has token', hastoken);

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
      $flare.auth
        .isAuthenticated()
        .then(function () {
          next({ name: 'index' });
        })
        .catch(function (err) {
          console.error(err);
          next();
        });
    } else {
      return next();
    }
  });

  // handle unauthorized requests
  window.BasicHttp.onUnauthorized = function () {
    var pending = appRouter.history.pending || {};
    var current = appRouter.history.current;
    if (current.name != 'login' && pending.name != 'login') {
      appRouter.push({ name: 'login' });
    }
  };
})(window.$flare);
