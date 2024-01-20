define([], function () {
  var helpers = {};

  helpers.VueLoader = function (vueFile) {
    return function (resolve) {
      return require([
        'vue!{{ .Data.Plugin.HttpApi.AssetPath "" }}/' + vueFile
      ], resolve);
    };
  };

  helpers.AssetPath = function (path) {
    return '{{ .Data.Plugin.HttpApi.AssetPath "" }}/' + path;
  };

  var routes = JSON.parse('{{ .Data.Routes }}');
  helpers.RoutePath = function (name) {
    var route;
    for (var i = 0; i < routes.length; i++) {
      if (routes[i].name === '{{ .Data.Plugin.Pkg }}.' + name) {
        route = routes[i];
        break;
      }
    }
    return route ? route.path : '{{ .Data.NotFoundPath }}';
  };

  return helpers;
});
