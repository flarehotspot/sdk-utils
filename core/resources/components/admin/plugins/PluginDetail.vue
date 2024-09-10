<template>
  <div>
    <div class="d-flex p-2">
      <h3>Plugin {{ data.Name }}</h3>
    </div>

    <hr />

    <p>{{ data }}</p>
  </div>
</template>

<script>
define(function () {
  return {
    template: template,
    data: function () {
      return {
        data: []
      };
    },
    mounted: function () {
      var self = this;
      $flare.http
        .get('<% .Helpers.UrlForRoute "admin:plugins:store:plugin"  %>')
        .then(function (data) {
          self.data = data;
          console.log(data);
        });
    },
    methods: {}
  };
});
</script>
