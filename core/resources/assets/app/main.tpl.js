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
      mounted: function () {
        var self = this;
        // handle unauthorized requests
        // window.BasicHttp.onUnauthorized = function () {
        //   var pending = self.$router.history.pending || {};
        //   var current = self.$router.history.current;
        //   if (current.name != 'login' && pending.name != 'login') {
        //     router.push({ name: 'login' });
        //   }
        // };

        // $flare._triggerReady();
      }
    });

    app.$mount('#app');

    $flare.app = app;
  });
})(window.$flare, window.Vue);
