/**
 * @file             : vue-lazy-load.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.ph>
 * Date              : Jan 25, 2024
 * Last Modified Date: Jan 25, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */
(function ($flare) {
  $flare.vueLazyLoad = function (vueFile) {
    return function (resolve) {
      return require(['vue!' + vueFile], resolve);
    };
  };
})(window.$flare);
