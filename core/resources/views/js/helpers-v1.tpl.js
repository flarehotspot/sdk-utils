define(['{{ .Data.CoreApi.HttpApi.AssetPath "services/http.js" }}'], function (
  http
) {
  var h = { http: http };

  h.VueLoader = function (vueFile) {
    return function (resolve) {
      return require([
        'vue!{{ .Data.Plugin.HttpApi.AssetPath "" }}/' + vueFile
      ], resolve);
    };
  };

  h.AssetPath = function (path) {
    return '{{ .Data.Plugin.HttpApi.AssetPath "" }}/' + path;
  };

  var routes = JSON.parse('{{ .Data.Routes }}');
  h.RoutePath = function (name) {
    var route;
    for (var i = 0; i < routes.length; i++) {
      if (routes[i].name === '{{ .Data.Plugin.Pkg }}.' + name) {
        route = routes[i];
        break;
      }
    }
    return route ? route.path : '{{ .Data.NotFoundPath }}';
  };

  return h;
});
