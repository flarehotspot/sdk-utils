<template>
  <div>
    <h1>Logs</h1>
    <button @click.prevent="clear">Clear Logs</button>
    <form>
      <div>
        <label>Per Page:</label>
        <select v-model="perPage" @change="navigate($event, currentPage)">
          <option value="10">10</option>
          <option value="50">50</option>
          <option value="100">100</option>
        </select>
      </div>
    </form>

    <table class="table table-striped">
      <thead>
        <tr>
          <th>Log</th>
          <th>Level</th>
          <th>Plugin</th>
          <th>File</th>
          <th>Time</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="log in logs">
          <td>
            <strong>{{ log.title }}</strong
            ><br />
            <p v-for="p in log.body">{{ p }}</p>
          </td>
          <td>{{ log.level }}</td>
          <td>{{ log.plugin }}</td>
          <td>{{ log.fullpath }}</td>
          <td>{{ log.datetime }}</td>
        </tr>
      </tbody>
    </table>

    <b-pagination
      v-model="currentPage"
      :per-page="perPage"
      :total-rows="count"
      @page-click="navigate"
    ></b-pagination>
  </div>
</template>

<script>
define(function () {
  return {
    template: template,
    data: function () {
      var self = this;
      return {
        selectedLevel: 'all',
        selectedPlugin: 'all',
        plugins: ['all'],
        searchFilter: self.$route.query.search || '',
        lines: parseInt(self.$route.query.lines || 50),
        currentPage: parseInt(self.$route.query.currentPage || 1),
        perPage: parseInt(self.$route.query.perPage || 50),
        count: 0,
        logs: []
      };
    },
    mounted: function () {
      var self = this;
      self.load();
    },
    methods: {
      load: function (event, currentPage) {
        var self = this;
        var params = {
          currentPage: currentPage || self.currentPage || 1,
          perPage: self.perPage
        };

        $flare.http
          .get('<% .Helpers.UrlForRoute "admin:logs:index" %>', params)
          .then(function (data) {
            console.log('Logs data:', data);
            self.logs = data.logs;
            self.count = data.count;
          });
      },
      navigate: function (event, currentPage) {
        var self = this;
        var params = {
          currentPage: currentPage || self.currentPage || 1,
          perPage: self.perPage
        };

        self.$router.replace({
          name: '<% .Helpers.VueRouteName "log-viewer" %>',
          query: params
        });

        self.load(event, currentPage);
      },
      clear: function (event) {
        var self = this;
        if (window.confirm('Are you sure you want to clear all logs?')) {
          $flare.http
            .post('<% .Helpers.UrlForRoute "admin:logs:clear" %>')
            .then(function () {
              self.count = 0;
              self.logs = [];
              self.currentPage = 1;
            })
            .catch(function (err) {
              console.error(err);
            });
        }
      }
    }
  };
});
</script>
