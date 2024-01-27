<template>
  <wrapped-component :flare-view="view"></wrapped-component>
</template>

<script>
(function (window) {
  var $flare = window.$flare;
  var compPath = 'vue!<% .Data.HttpComponentFullPath %>';
  var dataPath = '<% .Data.HttpDataFullPath %>';

  define([compPath], function (comp) {
    var viewData = { view: { loading: true, data: {}, errors: {} } };

    return {
      template: template,
      components: {
        WrappedComponent: comp
      },
      data: function () {
        return viewData;
      },
      mounted: function () {
        var params = this.$route.params;
        var path = substitutePathParams(dataPath, params);

        $flare.http
          .get(path)
          .then(function (data) {
            viewData.view.data = data;
          })
          .catch(function (res) {
            viewData.view.errors = res;
          })
          .finally(function () {
            viewData.view.loading = false;
          });
      }
    };
  });

  function substitutePathParams(path, params) {
    // Regular expression to match {param} in the path
    const paramRegex = /\{([^}]+)\}/g;

    // Replace each {param} with the corresponding value from the params object
    const substitutedPath = path.replace(paramRegex, function (_, paramName) {
      // If the param exists in the params object, use its value, otherwise, keep the original {param}
      return params.hasOwnProperty(paramName)
        ? params[paramName]
        : '{' + paramName + '}';
    });

    return substitutedPath;
  }
})(window);
</script>
