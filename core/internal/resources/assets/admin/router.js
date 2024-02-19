/**
 * @file             : router.js
 * @author           : Adones Pitogo <adones.pitogo@adopisoft.com>
 * Date              : Jan 19, 2024
 * Last Modified Date: Jan 22, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */

(function (window) {
  var $flare = window.$flare;
  var VueRouter = window.VueRouter;
  var routes = JSON.parse('<% .Data.Routes %>');
  // console.log(routes);
  routes = transformRoutes(routes);
  var router = new VueRouter({ routes: routes });

  router.beforeEach(function (to, _, next) {
    var hastoken = hasAuthToken();
    if (
      to.matched.some(function (route) {
        return route.meta.requireAuth;
      })
    ) {
      hastoken ? next() : next({ name: '<% .Data.LoginRouteName %>' });
    } else {
      next();
    }
  });

  // progress bar
  router.beforeResolve(function (_, to, next) {
    // If this isn't an initial page load.
    if (to.name) {
      // Start the route progress bar.
      NProgress.start();
    }
    next();
  });

  router.afterEach(function () {
    // Complete the animation of the route progress bar.
    NProgress.done();
  });

  function transformRoutes(routes) {
    var newRoutes = [];
    for (var i = 0; i < routes.length; i++) {
      var r = routes[i];
      var route = {};

      if (!r.redirect) {
        route = {
          name: r.name,
          path: r.path,
          component: $flare.vueLazyLoad(r.component),
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

  function hasAuthToken() {
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

  window.BasicHttp.onUnauthorized = function () {
    router.push({ name: '<% .Data.LoginRouteName %>' });
  };

  $flare.router = router;
})(window);
