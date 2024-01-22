/**
 * @file             : flare-view.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.com>
 * Date              : Jan 21, 2024
 * Last Modified Date: Jan 21, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */
(function (window) {
  var Vue = window.Vue;
  var $flare = window.$flare;
  var viewData = { view: { loading: false, data: {} } };

  Vue.component('FlareView', {
    template: '<router-view :flare-view="view"></router-view>',
    data: function () {
      return viewData;
    }
  });

})(window);
