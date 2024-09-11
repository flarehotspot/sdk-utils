<template>
  <div>
    <div class="d-flex p-2 align-items-center">
      <router-link
        class="mr-3"
        to='<% .Helpers.VueRoutePath "plugins-store" %>'
      >
        back
      </router-link>
      <h3>Plugin {{ data.Name }}</h3>
    </div>

    <hr />

    <div v-for="pr in data.Releases" :key="pr.Id">
      <a :href="pr.ZipFileUrl">
        <p>{{ pr.Major + '.' + pr.Minor + '.' + pr.Patch }}</p>
      </a>
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
        pluginId: parseInt(self.$route.query.id) || 1
      };
    },
    mounted: function () {
      var self = this;
      self.load();
    },
    methods: {
      load: function () {
        var self = this;
        console.log(self.$route.query);

        self.pluginId = parseInt(self.$route.query.id);

        console.log(self.pluginId);

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
            console.log(data);
          });
      }
    }
  };
});
</script>
