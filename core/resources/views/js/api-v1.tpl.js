define(['{{ .Data.CoreApi.HttpApi.AssetPath "services/http.js" }}'], function (
  http
) {
  var api = {};

  api.Http = http;

  // authentication
  api.Auth = {};
  api.Auth.IsAuthenticated = function () {
    return http.GetJson(
      '{{ .Helpers.UrlForMuxRoute "auth.is-authenticated" }}'
    );
  };
  api.Auth.Login = function (data) {
    return http.PostJson('{{ .Helpers.UrlForMuxRoute "auth.login" }}', data);
  };
  api.Auth.Logout = function () {
    return http.PostJson('{{ .Helpers.UrlForMuxRoute "auth.logout" }}');
  };

  // portal apis
  api.Portal = {}
  api.Portal.PortalItems = function(){
    return http.GetJson('{{ .Helpers.UrlForMuxRoute "portal.items" }}')
  }

  // admin apis
  api.Admin = {};
  api.Admin.NavMenu = function () {
    return http.GetJson('  {{ .Helpers.UrlForMuxRoute "admin.navs" }}');
  };

  return api;
});
