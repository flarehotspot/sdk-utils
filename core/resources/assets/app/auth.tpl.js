/**
 * @file             : auth.tpl.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.com>
 * Date              : Jan 20, 2024
 * Last Modified Date: Jan 20, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */
(function ($flare) {
  var http = window.BasicHttp;
  var auth = {};

  // authentication
  auth.isAuthenticated = function () {
    return http.GetJson(
      '<% .Helpers.UrlForMuxRoute "auth.is-authenticated" %>'
    );
  };
  auth.login = function (data) {
    return http.PostJson('<% .Helpers.UrlForMuxRoute "auth.login" %>', data);
  };
  auth.logout = function () {
    return http.PostJson('<% .Helpers.UrlForMuxRoute "auth.logout" %>');
  };

  auth.hasAuthToken = function () {
    var segmnts = document.cookie.split(';');
    var hastoken = false;
    for (var i = 0; i < segmnts.length; i++) {
      var seg = segmnts[i].split('=');
      if (seg[0].trim() === 'auth-token' && seg[1].length > 0) {
        hastoken = true;
        break;
      }
    }
    return hastoken;
  };

  $flare.auth = auth;
})(window.$flare);
