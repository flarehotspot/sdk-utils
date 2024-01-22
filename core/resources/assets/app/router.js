/**
 * @file             : router.js
 * @author           : Adones Pitogo <adones.pitogo@adopisoft.com>
 * Date              : Jan 19, 2024
 * Last Modified Date: Jan 22, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */

(function ($flare) {
  var VueRouter = window.VueRouter;
  var routes = JSON.parse('<% .Data.Routes %>');
  console.log(routes);
  routes = transformRoutes(routes);
  var router = new VueRouter({ routes: routes });

  router.beforeEach(function (to, _, next) {
    var hastoken = $flare.auth.hasAuthToken();

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

  function VueLazyLoad(vueFile) {
    return function (resolve) {
      return require(['vue!' + vueFile], resolve);
    };
  }

  function transformRoutes(routes) {
    var newRoutes = [];
    for (var i = 0; i < routes.length; i++) {
      var r = routes[i];
      var route = {};

      if (!r.redirect) {
        route = {
          name: r.name,
          path: r.path,
          component: VueLazyLoad(r.component),
          meta: r.meta
        };

        if (r.children) {
          route.children = transformRoutes(r.children);
        }
      } else {
        route = r;
      }

      newRoutes.push(route);
    }
    return newRoutes;
  }

  $flare.router = router;
})(window.$flare);
