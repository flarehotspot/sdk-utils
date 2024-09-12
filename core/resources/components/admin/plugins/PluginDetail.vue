<template>
  <div>
    <div class="d-flex p-2 align-items-center">
      <router-link
        class="mr-3 btn btn-secondary"
        to='<% .Helpers.VueRoutePath "plugins-store" %>'
      >
        back
      </router-link>
      <h3>{{ data.Name }}</h3>
    </div>

    <hr />

    <div class="d-flex" v-for="pr in data.Releases" :key="pr.Id">
      <p>{{ pr.Major + '.' + pr.Minor + '.' + pr.Patch }}</p>
      <div @click="installRelease($event, pr)" class="btn btn-primary">
        install
      </div>
    </div>
  </div>
</template>

<script>
define(function () {
  return {
    template: template,
    data: function () {
      var self = this;
      return {
        data: [],
        plugin: {},
        pluginId: parseInt(self.$route.query.id) || 1,
        def: {
          Src: 'store',
          GitURL: '',
          GitRef: ''
        }
      };
    },
    mounted: function () {
      var self = this;
      self.load();
    },
    methods: {
      load: function () {
        var self = this;

        self.pluginId = parseInt(self.$route.query.id);
        var params = {
          id: self.pluginId
        };

        $flare.http
          .get(
            '<% .Helpers.UrlForRoute "admin:plugins:store:plugin"  %>',
            params
          )
          .then(function (data) {
            self.data = data;

            self.plugin = {
              Id: self.data.Id,
              Name: self.data.Name,
              Package: self.data.Package
            };
          });
      },
      installRelease: function (e, pr) {
        e.preventDefault();
          
        var self = this;
        var params = {
          Src: 'store',
          StorePackage: self.plugin.Package,
          StoreZipFile: pr.ZipFileUrl
        };

        $flare.http
          .post('<% .Helpers.UrlForRoute "admin:plugins:install" %>', params)
          .then(function (response) {
            $flare.notify.success(`Plugin ${response.Name} installed`);
          });
      }
    }
  };
});
</script>
