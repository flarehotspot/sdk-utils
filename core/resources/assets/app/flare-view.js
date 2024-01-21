/**
 * @file             : flare-view.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.com>
 * Date              : Jan 21, 2024
 * Last Modified Date: Jan 21, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */
(function ($flare) {
  var viewData = { data: {} };

  Vue.component('flare-view', {
    template: '<router-view :data="data"></router-view>',
    data: function () {
      return viewData;
    },
    mounted: function () {
      var router = $flare.router;
      if (router.currentRoute.meta.data_path) {
        loadData(router.currentRoute);
      }

      router.afterEach(function (to, _) {
        loadData(to);
      });
    }
  });

  function loadData(route) {
    if (route.meta.data_path) {
      var data_path = route.meta.data_path;
      var params = route.params;
      var data_uri = substitutePathParams(data_path, params);
      BasicHttp.GetJson(data_uri).then(function (data) {
        viewData.data = data;
      });
    }
  }

  function substitutePathParams(path, params) {
    // Regular expression to match {param} in the path
    const paramRegex = /\{([^}]+)\}/g;

    // Replace each {param} with the corresponding value from the params object
    const substitutedPath = path.replace(paramRegex, (_, paramName) => {
      // If the param exists in the params object, use its value, otherwise, keep the original {param}
      return params.hasOwnProperty(paramName)
        ? params[paramName]
        : `{${paramName}}`;
    });

    return substitutedPath;
  }
})(window.$flare);
