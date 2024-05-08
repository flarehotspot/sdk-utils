/**
 * @file             : router.js
 * @author           : Adones Pitogo <adones.pitogo@adopisoft.com>
 * Date              : Jan 19, 2024
 * Last Modified Date: May 08, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
(function () {
  var VueRouter = window.VueRouter;
  var routesData = JSON.parse('<% .Data %>');
  var childRoutes = routesData.child_routes;

  childRoutes.push({
    path: routesData.dashboard_component.path,
    name: routesData.dashboard_component.name,
    component: $flare.vueLazyLoad(routesData.dashboard_component.component)
  });

  var routes = [
    {
      path: routesData.layout_component.path,
      name: routesData.layout_component.name,
      component: $flare.vueLazyLoad(routesData.layout_component.component),
      children: transformRoutes(childRoutes)
    },
    {
      path: routesData.login_component.path,
      name: routesData.login_component.name,
      component: $flare.vueLazyLoad(routesData.login_component.component)
    },
    {
      path: '*',
      redirect: {
        name: routesData.dashboard_component.name
      }
    }
  ];

  var router = new VueRouter({ routes: routes });
  $flare.router = router;

  router.beforeEach(function (to, _, next) {
    var hastoken = hasAuthToken();
    if (
      to.matched.some(function (route) {
        return route.meta.requireAuth;
      })
    ) {
      hastoken ? next() : next({ name: routesData.login_component.name });
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
          component:
            typeof r.component === 'string'
              ? $flare.vueLazyLoad(r.component)
              : r.component,
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
    console.log('error onUnauthorized');
    router.push({ name: routesData.login_component.name });
  };
})();
