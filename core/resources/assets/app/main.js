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

  define([], function () {
    var app = new Vue({
      router: $flare.router
      // store: $flare.store
    });

    app.$mount('#app');

    $flare.app = app;
  });
})(window);
