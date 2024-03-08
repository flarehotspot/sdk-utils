/**
 * @file             : main.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.com>
 * Date              : Jan 21, 2024
 * Last Modified Date: Feb 27, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>

 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
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

        $flare.events.on('session:connected', function (data) {
          $flare.notify.success(data.message);
        });

        $flare.events.on('session:disconnected', function (data) {
          $flare.notify.error(data.message);
        });
      }
    });

    app.$mount('#app');
    $flare.app = app;
  });
})(window);
