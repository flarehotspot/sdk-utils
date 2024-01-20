/**
 * Copyright 2021-2022 Flarego Technologies Corp. <business@flarego.ph>
 * @file             : domready-1.0.0.js
 * @author           : Adones Pitogo <adones.pitogo@adopisoft.com>
 * Date              : Dec 27, 2022
 * Last Modified Date: Jan 20, 2024
 */
window.DomReady =
  window.DomReady ||
  function (f) {
    window._domready_callbacks = window._domready_callbacks || [];
    window._domready_callbacks.push(f);
    if (window._domready_callbacks.length > 1) return;

    var callbacks = window._domready_callbacks;

    function readyCallback() {
      for (var i = 0; i < callbacks.length; i++) {
        var cb = callbacks[i];
        cb();
      }
    }

    function checkReady() {
      if (/in/.test(document.readyState)) {
        setTimeout(checkReady, 10); // check every 10ms
      } else {
        readyCallback();
      }
    }

    checkReady();
  };
