/**
 * @file             : require-config.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.com>
 * Date              : Jan 22, 2024
 * Last Modified Date: Jan 22, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */
var require = {
  paths: {
    vue: '<% .Helpers.AssetPath "libs/requirejs-vue-1.1.5.min" %>',
    image: '<% .Helpers.AssetPath "libs/requirejs-image-0.2.2.min" %>',
    text: '<% .Helpers.AssetPath "libs/requirejs-text-2.0.5.min" %>'
  },
  config: {
    vue: {
      css: 'inject',
      templateVar: 'template'
    }
  }
};
