/**
 * @file             : events.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.ph>
 * Date              : Jan 25, 2024
 * Last Modified Date: Feb 27, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>

 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

(function ($flare) {
  var sse = new EventSource('<% .Data %>');

  window.onbeforeunload = function () {
    sse.close();
  };

  $flare.events = {
    on: function (event, callback) {
      sse.addEventListener(event, function (res) {
        var data = JSON.parse(res.data);
        callback(data);
      });
    },
    off: function (event, callback) {
      sse.removeEventListener(event, callback);
    }
  };
})(window.$flare);
