<template>
  <div>
    <div class="d-flex p-2 align-items-center">
      <router-link
        class="mr-3 btn btn-secondary"
        to='<% .Helpers.VueRoutePath "plugins-store" %>'
      >
        back to store
      </router-link>
    </div>

    <hr />

    <h3>{{ data.Name }}</h3>

    <div
      class="btn btn-primary w-100"
      @click="installRelease($event, data.Releases[0])"
    >
      Install Plugin
    </div>

    <hr />
    <h4>Description</h4>
    <hr />

    <h4>Other Versions</h4>
    <table class="table table-bordered table-striped">
      <thead>
        <tr>
          <th>Version</th>
          <th>Option</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="pr in data.Releases" :key="pr.Id">
          <td>{{ pr.Major + '.' + pr.Minor + '.' + pr.Patch }}</td>
          <td>
            <button
              type="button"
              class="btn btn-secondary"
              @click="installRelease($event, pr)"
            >
              Install
            </button>
          </td>
        </tr>
      </tbody>
    </table>
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
            $flare.notify.success(`${response.Name} installed`);
          });
      }
    }
  };
});
</script>
