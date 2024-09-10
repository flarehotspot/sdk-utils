<template>
  <div>
    <div class="d-flex p-2">
      <h3>Flare Plugins Store</h3>
      <p>s</p>
    </div>

    <hr />

    <div v-for="p in plugins">
      <button class="border" @click="viewPlugin(p.Id)">
        <div class="">
          <p>{{ p.Name }}</p>
          <p>{{ p.Package }}</p>
        </div>
      </button>
    </div>
  </div>
</template>

<script>
define(function () {
  return {
    template: template,
    data: function () {
      return {
        plugins: []
      };
    },
    mounted: function () {
      var self = this;
      $flare.http
        .get('<% .Helpers.UrlForRoute "admin:plugins:store:index"  %>')
        .then(function (plugins) {
          self.plugins = plugins;
          console.log(plugins);
        });
    },
    methods: {
      viewPlugin: function (pluginId) {
        var self = this;

        var pluginData = {
          id: pluginId
        };

        self.$router.replace({
          name: '<% .Helpers.VueRouteName "plugin" %>',
          query: pluginData
        });
      }
    }
  };
});
</script>
