<template>
  <div>
    <h1>System Updates</h1>

    <hr />

    <div class="" v-if="shouldUpdate(current, latest)">
      <h1>Update Available</h1>
      <div class="btn btn-primary">Install Update</div>
    </div>
    <div v-else>
      <h1>Latest</h1>
    </div>

    <p>Latest Version: {{ stringifyVersion(latest) }}</p>
    <p>Installed Version: {{ stringifyVersion(current) }}</p>

    <div class="btn btn-secondary" @click="fetchLatest">Check Updates</div>
  </div>
</template>

<script>
define(function () {
  return {
    template: template,
    data: function () {
      return {
        latest: {},
        current: {},
        isUpToDate: false,
        localCoreFilesPath: '',
        localArchBinFilePath: ''
      };
    },
    mounted: function () {
      var self = this;
      self.fetchLatest();
      self.getCurrent();
    },
    methods: {
      fetchLatest: async function () {
        var self = this;

        await $flare.http
          .get('<% .Helpers.UrlForRoute "admin:core:fetch" %>')
          .then(function (response) {
            console.log(response);
            self.latest = response;
          })
          .catch(function (error) {
            console.log(error);
          });
      },
      getCurrent: async function () {
        var self = this;

        await $flare.http
          .get('<% .Helpers.UrlForRoute "admin:core:current" %>')
          .then(function (response) {
            console.log(response);
            self.current = response;
          })
          .catch(function (error) {
            console.log(error);
          });
      },
      stringifyVersion: function (version) {
        return 'v' + version.Major + '.' + version.Minor + '.' + version.Patch;
      },
      shouldUpdate: function (current, latest) {
        if (current.Major < latest.Major) {
          return true;
        }
        if (current.Minor < latest.Minor) {
          return true;
        }
        if (current.Patch < latest.Patch) {
          return true;
        }
        return false;
      }
    }
  };
});
</script>
