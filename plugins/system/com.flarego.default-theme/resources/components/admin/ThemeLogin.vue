<template>
  <div class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3">
        <flare-form @submit.prevent="login" heading="Admin Login">
          <flare-input-field label="Username" type="text" v-model="username">
          </flare-input-field>

          <flare-input-field
            label="Password"
            type="password"
            v-model="password"
          >
          </flare-input-field>

          <flare-button
            type="submit"
            :disabled="loading || !username || !password"
          >
            <span v-if="!loading">Login</span>
            <span v-else>Logging in...</span>
          </flare-button>
        </flare-form>
      </div>
    </div>
  </div>
</template>

<script>
(function ($flare) {
  define(function () {
    var comp = { template: template };

    comp.data = function () {
      return {
        loading: false,
        username: "",
        password: ""
      };
    };

    comp.methods = {
      login: function () {
        var self = this;

        self.loading = true;
        $flare.http
          .post('<% .Helpers.UrlForRoute "auth.login" %>', {
            username: self.username,
            password: self.password
          })
          .then(function () {
            self.$router.push({
              name: '<% .Helpers.VueRouteName "dashboard" %>'
            });
          })
          .finally(function () {
            self.loading = false;
          });
      }
    };

    return comp;
  });
})(window.$flare);
</script>
