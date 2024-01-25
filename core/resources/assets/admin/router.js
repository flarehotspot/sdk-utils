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
    var hastoken = hasAuthToken();
    console.log('Has token: ', hastoken);
    console.log('LoginRouteName: ', '<% .Data.LoginRouteName %>');

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
  router.beforeResolve(function (from, to, next) {
    // If this isn't an initial page load.
    if (to.name) {
      // Start the route progress bar.
      NProgress.start();
    }
    next();
  });

  router.afterEach(function (to, from) {
    // Complete the animation of the route progress bar.
    NProgress.done();
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

  $flare.router = router;
})(window.$flare);
