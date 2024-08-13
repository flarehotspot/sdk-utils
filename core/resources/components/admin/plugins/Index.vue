<template>
  <div>
    <h1>Plugins</h1>
    <hr />

    <router-link class="btn btn-primary" to='<% .Helpers.VueRoutePath "plugins-new" %>'>Install Plugin</router-link>

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
              <button type="button" class="btn btn-danger">Uninstall</button>
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
    }
  };
});
</script>
