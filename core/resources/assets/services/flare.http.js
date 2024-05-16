/**
 * @file             : vue-http.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.com>
 * Date              : Jan 22, 2024
 * Last Modified Date: Feb 27, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>

 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
(function ($flare) {
  var basicHttp = window.BasicHttp;
  var http = {};
  var rootres = '$vue_response';

  http.get = function (url, params) {
    var opts = {
      headers: { Accept: 'application/json' }
    };

    return basicHttp
      .GetJson(url, params, opts)
      .then(function (data) {
        return parseRespones(data);
      })
      .catch(function (res) {
        return Promise.reject(parseRespones(res));
      });
  };

  http.post = function (url, params) {
    var opts = {
      headers: { 'Content-Type': 'application/json' }
    };

    return basicHttp
      .PostJson(url, params, opts)
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
      var redirect = $res.redirect;
      $flare.router.push({
        name: redirect.routename,
        params: redirect.params,
        query: redirect.query
      });
      return {};
    } else if ($res.data !== undefined) {
      return $res.data;
    } else {
      invalidResponse(res);
      return res;
    }
  }

  $flare.http = http;
})(window.$flare);
