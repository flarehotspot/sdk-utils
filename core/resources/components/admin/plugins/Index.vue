<template>
  <div>
    <h1>Plugins</h1>
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
      $flare.http
        .get('<% .Helpers.UrlForRoute "admin:plugins:index"  %>')
        .then(function (res) {
          console.log(res);
        });
    }
  };
});
</script>
