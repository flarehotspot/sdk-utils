<template>
  <div>
    <h1>Flarehotspot Plugins Store</h1>
    <hr />
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
