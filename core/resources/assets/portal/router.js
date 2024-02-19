/**
 * @file             : router.js
 * @author           : Adones Pitogo <adones.pitogo@adopisoft.com>
 * Date              : Jan 19, 2024
 * Last Modified Date: Jan 25, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */

(function ($flare) {
  var VueRouter = window.VueRouter;
  var routes = JSON.parse('<% .Data.Routes %>');
  console.log(routes);
  routes = transformRoutes(routes);
  var router = new VueRouter({ routes: routes });
  $flare.router = router;

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
})(window.$flare);
