/**
 * @file             : main.tpl.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.com>
 * Date              : Jan 21, 2024
 * Last Modified Date: Jan 21, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */

// var vueNotifsPath = '<% .Helpers.AssetPath "libs/vue-notification.js" %>'

(function ($flare, Vue) {
  define([], function () {
    // console.log(VueNotif);

    var app = new Vue({
      router: $flare.router,
    });

    app.$mount('#app');

    $flare.app = app;
  });
})(window.$flare, window.Vue);
