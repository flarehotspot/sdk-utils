<template>
  <div>
    <h1>Installed Plugins</h1>
    <hr />

    <div class="mb-2">
      <router-link
        class="btn btn-primary"
        to='<% .Helpers.VueRoutePath "plugins-new" %>'
        >Install a Plugin</router-link
      >

      <router-link
        class="btn btn-primary"
        to='<% .Helpers.VueRoutePath "plugins-store" %>'
        >Visit Flare Plugins</router-link
      >
    </div>

    <table class="table table-bordered table-striped">
      <thead>
        <tr>
          <th>Name</th>
          <th>Description</th>
          <th>Version</th>
          <th>Options</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="p in plugins">
          <td>{{ p.Info.Name }}</td>
          <td>{{ p.Info.Description }}</td>
          <td>{{ p.Info.Version }}</td>
          <td>
            <button
              type="button"
              class="btn btn-danger"
              v-on:click="uninstall(p.Info.Package)"
            >
              Uninstall
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
      return {
        plugins: []
      };
    },
    mounted: function () {
      var self = this;
      $flare.http
        .get('<% .Helpers.UrlForRoute "admin:plugins:index"  %>')
        .then(function (plugins) {
          self.plugins = plugins;
        });
    },
    methods: {
      uninstall: function (pkg) {
        var self = this;
        var yes = confirm('Are you sure you want to uninstall this plugin?');
        if (!yes) {
          return;
        }

        $flare.http
          .post('<% .Helpers.UrlForRoute "admin:plugins:uninstall" %>', {
            pkg: pkg
          })
          .then(function () {
            console.log('Uninstalled plugin: ' + pkg);
          });
      }
    }
  };
});
</script>
