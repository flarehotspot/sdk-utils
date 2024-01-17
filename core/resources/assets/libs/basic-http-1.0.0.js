/**
 * Copyright 2021-2022 Flarego Technologies Corp. <business@flarego.ph>
 * @file             : http.js
 * @author           : Adones Pitogo <pitogo.adones@gmail.com>
 * Date              : Nov 29, 2022
 * Last Modified Date: Nov 29, 2022
 */

var BasicHttp = (function () {
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

  function handleError(errorCb, client) {
    var body = client.response || client.responseText;
    var error = 'Something went wrong';

    if (
      (client.status === 401 || client.status === 403) &&
      typeof http.onUnauthorized === 'function'
    ) {
      http.onUnauthorized({ error: error });
    }

    if (body) {
      try {
        var data = JSON.parse(body);
        error = data.error || data.message || error;
        errorCb({ error: error });
      } catch (e) {
        console.error(e);
        errorCb({ error: error });
      }
    } else {
      errorCb({ error: error });
    }
  }

  function Ajax(opts) {
    var noop = function () {
      console.log('http callback not defined');
    };

    var method = (opts.method || 'GET').toUpperCase();
    var url = opts.url;
    var data = opts.data || {};
    var successCb = opts.success || noop;
    var errorCb = opts.error || noop;

    var client = httpClient();

    client.onreadystatechange = function () {
      if (client.readyState === 4) {
        if (client.status >= 200 && client.status < 400) {
          try {
            var data = JSON.parse(client.responseText);
            successCb(data);
          } catch (e) {
            console.log(e);
            handleError(errorCb, client);
          }
        } else {
          handleError(errorCb, client);
        }
      }
    };

    if (method === 'GET') {
      // prevent ajax caching
      data.cache_bust = Math.random().toString().replace('.', '');
      try {
        url += url.indexOf('?') > -1 ? '&' : '?';
        url += serialize(data);
        client.open(method, url, true);
        client.send();
      } catch (e) {
        handleError(errorCb, client);
      }
    } else if (method === 'POST') {
      try {
        client.open(method, url, true);
        client.setRequestHeader(
          'Content-Type',
          'application/x-www-form-urlencoded'
        );
        var params = serialize(data);
        client.send(params);
      } catch (e) {
        handleError(errorCb, client);
      }
    } else {
      console.error('Unsupported method:', method);
    }
  }

  var http = {};

  http.GetJson = function (url, data) {
    return new Promise(function (resolve, reject) {
      try {
        Ajax({
          url: url,
          data: data,
          success: resolve,
          error: reject
        });
      } catch (e) {
        console.error(e);
        cb(e);
      }
    });
  };

  http.PostJson = function (url, data) {
    return new Promise(function (resolve, reject) {
      try {
        Ajax({
          method: 'POST',
          url: url,
          data: data,
          success: function (data) {
            resolve(data);
          },
          error: function (e) {
            reject(e);
          }
        });
      } catch (e) {
        callback(e);
      }
    });
  };

  http.onUnauthorized = function (err) {
    console.error('Unauthorezed Error!', err);
  };

  return http;
})();
