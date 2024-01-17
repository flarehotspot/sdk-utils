define(['{{ .Data.CoreApi.HttpApi.AssetPath "services/http.js" }}'], function (
  http
) {
  var api = {};

  api.Http = http;

  // authentication
  api.Auth = {};
  api.Auth.IsAuthenticated = function () {
    return http.getJson(
      '{{ .Helpers.UrlForMuxRoute "auth.is-authenticated" }}'
    );
  };
  api.Auth.Login = function (data) {
    return http.postJson('{{ .Helpers.UrlForMuxRoute "auth.login" }}', data);
  };
  api.Auth.Logout = function () {
    return http.postJson('{{ .Helpers.UrlForMuxRoute "auth.logout" }}');
  };

  // portal apis
  api.Portal = {}
  api.Portal.PortalItems = function(){
    return http.getJson('{{ .Helpers.UrlForMuxRoute "portal.items" }}')
  }

  // admin apis
  api.Admin = {};
  api.Admin.NavMenu = function () {
    return http.getJson('  {{ .Helpers.UrlForMuxRoute "admin.navs" }}');
  };

  return api;
});
