/**
 * @file             : navigator.tpl.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.com>
 * Date              : Jan 20, 2024
 * Last Modified Date: Jan 20, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */
(function ($flare) {
  var Vue = window.Vue;
  var router = $flare.router;

  var viewData = { data: {} };

  Vue.component('flare-view', {
    template: '<router-view :data="data"></router-view>',
    data: function () {
      return viewData;
    },
    mounted: function () {
      var router = $flare.router;
      if (router.currentRoute.meta.data_path) {
        var data_path = router.currentRoute.meta.data_path;
        BasicHttp.GetJson(data_path).then((data) => {
          viewData.data = data;
        });
      }
    }
  });

  Vue.component('flare-link', {
    props: ['to'],
    template: '<a :href="href"><slot></slot></a>',
    data: function () {
      return {
        href: ''
      };
    },
    mounted: function () {
      var self = this;
      var el = self.$el;
      var to = self.to;

      var record = router.resolve({
        name: 'com.flarego.basic-system-account.accounts-list'
      });
      var data_path = record.route.meta.data_path;
      console.log(record.route);
      // BasicHttp.GetJson(data_path).then((data) =>
      //   console.log('component data: ', data)
      // );

      el.addEventListener('click', function (e) {
        e.preventDefault();
        router.push({ name: 'login', params: { sample: 'data' } });
      });
    }
  });
})(window.$flare);
