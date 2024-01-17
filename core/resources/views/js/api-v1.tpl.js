define([], function () {
  var api = {};

  api.Http = window.BasicHttp;

  // authentication
  api.Auth = {};
  api.Auth.IsAuthenticated = function () {
    return api.Http.GetJson(
      '{{ .Helpers.UrlForMuxRoute "auth.is-authenticated" }}'
    );
  };
  api.Auth.Login = function (data) {
    return api.Http.PostJson(
      '{{ .Helpers.UrlForMuxRoute "auth.login" }}',
      data
    );
  };
  api.Auth.Logout = function () {
    return api.Http.PostJson('{{ .Helpers.UrlForMuxRoute "auth.logout" }}');
  };

  // portal apis
  api.Portal = {};
  api.Portal.PortalItems = function () {
    return api.Http.GetJson('{{ .Helpers.UrlForMuxRoute "portal.items" }}');
  };

  // admin apis
  api.Admin = {};
  api.Admin.NavMenu = function () {
    return api.Http.GetJson('  {{ .Helpers.UrlForMuxRoute "admin.navs" }}');
  };

  return api;
});
