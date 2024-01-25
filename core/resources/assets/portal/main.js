/**
 * @file             : main.tpl.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.com>
 * Date              : Jan 21, 2024
 * Last Modified Date: Jan 21, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */

(function (window) {
  var $flare = window.$flare || {};
  var Vue = window.Vue;
  var viewData = { view: { loading: false, data: {} } };

  define([], function () {
    var app = new Vue({
      router: $flare.router,
      data: function () {
        return viewData;
      },
      mounted: function () {
        viewData.view.loading = false;
      }
    });

    app.$mount('#app');
    $flare.app = app;
  });
})(window);
