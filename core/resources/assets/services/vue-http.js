/**
 * @file             : vue-http.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.com>
 * Date              : Jan 22, 2024
 * Last Modified Date: Jan 22, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */
(function ($flare) {
  var http = window.BasicHttp;
  var vueHttp = {};
  var rootres = '$$response';
  $flare.http = vueHttp;

  vueHttp.get = function (url, params) {
    return http
      .GetJson(url, params)
      .then(function (data) {
        return parseRespones(data);
      })
      .catch(function (res) {
        return Promise.reject(parseRespones(res));
      });
  };

  vueHttp.post = function (url, params) {
    return http
      .PostJson(url, params)
      .then(function (data) {
        return parseRespones(data);
      })
      .catch(function (res) {
        return Promise.reject(parseRespones(res));
      });
  };

  function invalidResponse(res) {
    var err = res.error || res.message || 'Invalid server response';
    $flare.notify.error(err);
  }

  function parseRespones(res) {
    var $res = res[rootres];
    if (!$res) {
      invalidResponse(res);
      return res;
    }

    if ($res.flash) {
      var toastFn = $flare.notify[$res.flash.type];
      if (toastFn) {
        toastFn($res.flash.msg);
      } else {
        console.error('Invalid flash type:', $res.flash.type);
        $flare.notify.info($res.flash.msg);
      }
    }

    if ($res.redirect) {
      $flare.router.push({
        name: $res.routename,
        params: $res.params,
        query: $res.query
      });
      return {};
    } else if ($res.data !== undefined) {
      return $res.data;
    } else {
      invalidResponse(res);
      return res;
    }
  }
})(window.$flare);
