/**
 * @file             : utils.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.ph>
 * Date              : Feb 27, 2024
 * Last Modified Date: Feb 27, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>

 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
(function ($flare) {
  $flare.utils = {
    strHasChar: function (str, ch) {
      for (var i = 0; i < str.length; i++) {
        if (str[i] === ch) {
          return true;
        }
      }
    },

    attachQueryParams: function (path, query) {
      var keys = Object.keys(query);
      if (keys.length > 0) {
        var queryarr = [];
        for (var i = 0; i < keys.length; i++) {
          if (Object.hasOwnProperty.call(query, keys[i])) {
            queryarr.push(keys[i] + '=' + query[keys[i]]);
          }
        }
        if (!$flare.utils.strHasChar(path, '?')) {
          path += '?' + queryarr.join('&');
        } else {
          path += '&' + queryarr.join('&');
        }
      }

      return path;
    },

    vuePathToMuxPath: function (path, params) {
      // Regular expression to match {param} in the path
      const paramRegex = /\{([^}]+)\}/g;

      // Replace each {param} with the corresponding value from the params object
      const substitutedPath = path.replace(paramRegex, function (_, paramName) {
        // If the param exists in the params object, use its value, otherwise, keep the original {param}
        return params.hasOwnProperty(paramName)
          ? params[paramName]
          : '{' + paramName + '}';
      });

      return substitutedPath;
    }
  };
})(window.$flare);
