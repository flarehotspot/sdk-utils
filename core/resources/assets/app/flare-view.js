/**
 * @file             : flare-view.js
 * @author           : Adones Pitogo <pitogo.adones@flarego.com>
 * Date              : Jan 21, 2024
 * Last Modified Date: Jan 21, 2024
 * Copyright 2021-2024 Flarego Technologies Corp. <business@flarego.ph>
 */
(function (window) {
  var Vue = window.Vue;
  var Vuex = window.Vuex;
  var $flare = window.$flare;

  var $store = new Vuex.Store({
    state: {
      $view: {
        loading: true,
        data: {}
      }
    },
    mutations: {
      setLoading: function (state, loading) {
        state.$view.loading = loading;
      },
      setData: function (state, data) {
        state.$view.data = data;
      }
    }
  });

  Vue.component('flare-view', {
    template: '<router-view></router-view>',
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

      $flare.http
        .get(data_uri)
        .then(function (data) {
          console.log(data);
          $store.commit('setLoading', false);
          $store.commit('setData', data);
        })
        .finally(function () {
          $store.commit('setLoading', false);
        });
    } else {
      $store.commit('setData', {});
    }
  }

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

  $flare.view = $store.state.$view;
})(window);
