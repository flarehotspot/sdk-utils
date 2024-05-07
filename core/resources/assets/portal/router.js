/**
 * @file             : router.js
 * @author           : Adones Pitogo <adones.pitogo@adopisoft.com>
 * Date              : Jan 19, 2024
 * Last Modified Date: Feb 27, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
(function ($flare) {
  var VueRouter = window.VueRouter;
  var childRoutes = JSON.parse('<% .Data.ChildRoutes %>');
  // routes = transformRoutes(routes);

  var portalThemeComponent = {
    template: '<theme-layout></theme-layout>',
    components: {
      'theme-layout': $flare.vueLazyLoad('<% .Data.ThemeComponent.Component %>')
    },
    mounted: function () {
      $flare.http
        .get('<% .Helpers.UrlForRoute "portal.items" %>')
        .then(function (data) {
          console.log('nav items', data);
        });
    }
  };

  var routes = [
    {
      path: '<% .Data.ThemeComponent.Path %>',
      name: '<% .Data.ThemeComponent.Name %>',
      component: portalThemeComponent,
      children: transformRoutes(childRoutes)
    }
  ];

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
