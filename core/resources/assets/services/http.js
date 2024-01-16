/**
 * Copyright 2021-2022 Flarego Technologies Corp. <business@flarego.ph>
 * @file             : http.js
 * @author           : Adones Pitogo <pitogo.adones@gmail.com>
 * Date              : Nov 29, 2022
 * Last Modified Date: Nov 29, 2022
 */
(function () {
  var http = (function () {
    function httpClient() {
      return window.ActiveXObject
        ? new window.ActiveXObject('Microsoft.XMLHTTP')
        : new window.XMLHttpRequest();
    }

    function serialize(obj) {
      var str = [];

      for (var p in obj) {
        if (Object.prototype.hasOwnProperty.call(obj, p)) {
          str.push(encodeURIComponent(p) + '=' + encodeURIComponent(obj[p]));
        }
      }
      return str.join('&');
    }

    function Ajax(opts) {
      var method = (opts.method || 'GET').toUpperCase();
      var url = opts.url;
      var data = opts.data || {};
      var successCb =
        opts.success ||
        function () {
          console.log('http success callback not defined');
        };
      var errorCb =
        opts.error ||
        function () {
          console.log('http error callback not defined');
        };

      var http = httpClient();

      http.onreadystatechange = function () {
        if (http.readyState === 4) {
          if (http.status >= 200 && http.status < 400) {
            try {
              var json = JSON.parse(http.responseText);
              successCb(json);
            } catch (e) {
              errorCb(e);
            }
          } else {
            errorCb(http);
          }
        }
      };

      // prevent ajax caching
      if (method === 'GET') {
        var cache_bust = Math.random().toString().replace('.', '');
        url += url.indexOf('?') > -1 ? '&' : '?';
        url += serialize({ cache_bust: cache_bust });
      }

      http.open(method, url, true);
      http.setRequestHeader('Accept', 'application/json');

      if (method === 'POST') {
        try {
          // Send the proper header information along with the request
          http.setRequestHeader('Content-type', 'application/json');
          var params = JSON.stringify(data);
          http.send(params);
        } catch (e) {
          errorCb(e);
        }
      } else {
        http.send();
      }
    }

    var http = {};

    http.getJson = function (url, cb) {
      try {
        Ajax({
          url: url,
          success: function (data) {
            try {
              cb(null, data);
            } catch (e) {
              console.error('Error in BasicHttp#get callback:', e);
            }
          },
          error: cb
        });
      } catch (e) {
        console.error(e);
        cb(e);
      }
    };

    http.postJson = function (url, data, cb) {
      var callback = cb;
      if (typeof data === 'function') {
        callback = data;
      }
      try {
        data.tmp_client_id = http.tmp_client_id;
        Ajax({
          url: url,
          method: 'POST',
          data: typeof data === 'function' ? {} : data,
          success: function (data) {
            callback(null, data);
          },
          error: function (e) {
            callback(e);
          }
        });
      } catch (e) {
        callback(e);
      }
    };

    return http;
  })();

  define(function () {
    return http;
  });
})();
