<template>
  <div>
    <div class="d-flex flex-row p-2 align-items-center">
      <router-link
        to='<% .Helpers.VueRoutePath "plugins-index" %>'
        class="btn btn-secondary mr-2"
      >
        back
      </router-link>
      <h3>Flare Plugins</h3>
      <div class="w-100"></div>
      <p>search</p>
    </div>
    <hr />

    <div class="row">
      <div v-for="p in plugins" :key="p.Id" class="col-md-4 col-3 mb-2">
        <div
          class="bg-light border rounded rounded-4 p-2 cursor-pointer mb-2"
          @click="viewPlugin(p.Id)"
          role="button"
        >
          <div class="d-flex">
            <h3 class="w-100">{{ p.Info.Name }}</h3>

            <p
              class="text-light bg-success p-2 rounded-pill"
              v-if="p.IsInstalled"
            >
              installed
            </p>
          </div>
          <p>{{ p.Info.Package }}</p>
        </div>
      </div>
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
      self.load();
    },
    methods: {
      load: function () {
        var self = this;

        $flare.http
          .get('<% .Helpers.UrlForRoute "admin:plugins:store:index" %>')
          .then(function (pluginsData) {
            self.plugins = pluginsData;
          });
      },
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
