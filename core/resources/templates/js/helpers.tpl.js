define(function () {
  var routes = JSON.parse(`{{ .Data.Routes }}`);
  var h = {};

  h.vueLoader = function (vueFile) {
    return function (resolve) {
      return require([
        'vue!' + '{{ .Data.Plugin.AssetPath "" }}/' + vueFile
      ], resolve);
    };
  };

  h.routePath = function (name) {
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
