/**
 * @file             : main.tpl.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.com>
 * Date              : Jan 21, 2024
 * Last Modified Date: Jan 21, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */

(function (window) {
  var $flare = window.$flare || {};
  var router = $flare.router;
  var http = $flare.http;
  var Vue = window.Vue;
  var viewData = { view: { loading: false, data: {} } };

  // router.beforeEach(function (to, _, next) {
  //   console.log('To Route: ', to);
  //   next();
  // });

  define([], function () {
    var app = new Vue({
      router: $flare.router,
      data: function () {
        return viewData;
      },
      mounted: function () {
        // var themeRoute = router.resolve({ name: 'layout' });
        // var dataPath = themeRoute.route.meta.data_path;
        // viewData.view.loading = true;

        // http
        //   .get(dataPath)
        //   .then(function (data) {
        //     console.log('Layout Data', data);
        //     viewData.view.data = data;
        //   })
        //   .catch(function (err) {
        //     console.error(err);
        //     $flare.notify.error(
        //       err.error || err.message || 'Error fetching theme layout data'
        //     );
        //   })
        //   .finally(function () {
        //     viewData.view.loading = false;
        //   });
      }
    });

    app.$mount('#app');

    $flare.app = app;
  });
})(window);
