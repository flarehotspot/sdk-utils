define(function () {
  var routes = JSON.parse(`{{ .Data.Routes }}`);
  var h = {};

  h.plugin = {
    name: "{{ .Data.Plugin.Name }}",
    package: "{{ .Data.Plugin.Pkg }}"
  };

  h.vueLoader = function (vueFile) {
    if (vueFile.charAt(0) === "/") {
      vueFile = vueFile.substring(1);
    }

    return function (resolve) {
      return require(["vue!" + "{{ .Data.Plugin.Pkg }}/" + vueFile], resolve);
    };
  };

  h.routePath = function (name) {
    var route;
    for (var i = 0; i < routes.length; i++) {
      if (routes[i].name === "{{ .Data.Plugin.Pkg }}." + name) {
        route = routes[i];
        break;
      }
    }
    return route ? route.path : "";
  };

  return h;
});
