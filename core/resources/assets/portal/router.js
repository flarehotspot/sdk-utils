/**
 * @file             : router.js
 * @author           : Adones Pitogo <adones.pitogo@adopisoft.com>
 * Date              : Jan 19, 2024
 * Last Modified Date: May 07, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
(function ($flare) {
  var VueRouter = window.VueRouter;
  var routesData = JSON.parse('<% .Data %>');
  var childRoutes = routesData.child_routes;
  var reloadListener;
  var portalIndexComponent = {
    template: '<theme-index :data="data"></theme-index>',
    components: {
      'theme-index': $flare.vueLazyLoad(routesData.index_component.component)
    },
    data: function () {
      return {
        data: {
          loading: true,
          portalItems: []
        }
      };
    },
    mounted: function () {
      var self = this;
      self.load();

      reloadListener = $flare.events.on(
        'portal:items:reload',
        function (items) {
          self.data.portalItems = items;
        }
      );
    },
    beforeDestroy: function () {
      if (reloadListener) {
        $flare.events.off('portal:items:reload', reloadListener);
      }
    },
    methods: {
      load: function () {
        var self = this;
        $flare.http
          .get('<% .Helpers.UrlForRoute "portal:navs:items" %>')
          .then(function (data) {
            self.data.portalItems = data;
          })
          .finally(function () {
            self.data.loading = false;
          });
      }
    }
  };

  childRoutes.push({
    path: routesData.index_component.path,
    name: routesData.index_component.name,
    component: portalIndexComponent
  });

  var routes = [
    {
      path: routesData.layout_component.path,
      name: routesData.layout_component.name,
      component: $flare.vueLazyLoad(routesData.layout_component.component),
      children: transformRoutes(childRoutes)
    },
    {
      path: '*',
      redirect: {
        name: routesData.index_component.name
      }
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
})(window.$flare);
