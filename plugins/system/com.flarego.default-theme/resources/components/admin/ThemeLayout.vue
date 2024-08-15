<template>
  <div id="theme-layout">
    <b-navbar toggleable="lg" type="dark" variant="dark">
      <b-navbar-brand href='#<% .Helpers.VueRoutePath "dashboard" %>'>
        Dashboard
      </b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item-dropdown :text="list.label" right v-for="list in navs">
            <b-dropdown-item
              :href="'#' + nav.route_path"
              v-for="nav in list.items"
              >{{ nav.label }}</b-dropdown-item
            >
          </b-nav-item-dropdown>
        </b-navbar-nav>

        <b-navbar-nav class="ml-auto">
          <b-nav-item-dropdown right>
            <!-- Using 'button-content' slot -->
            <template #button-content>
              <em>User</em>
            </template>
            <b-dropdown-item href="#">Profile</b-dropdown-item>
            <b-dropdown-item href="#" @click="logout">Sign Out</b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>

    <div class="container-fluid">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
(function ($flare) {
  define(function () {
    var comp = {
      template: template,
      data: function () {
        return {
          navs: []
        };
      }
    };

    comp.mounted = function () {
      var self = this;
      $flare.http
        .get('<% .Helpers.UrlForRoute "admin.navs" %>')
        .then(function (navs) {
          self.navs = navs;
        });
    };

    comp.methods = {
      logout: function (e) {
        var self = this;
        console.log(e);
        $flare.http
          .post('<% .Helpers.UrlForRoute "auth.logout" %>')
          .then(function () {
            console.log("loggedout");
            self.$router.push({ name: '<% .Helpers.VueRouteName "login" %>' });
          })
          .catch(function (res) {
            console.error(res);
          });
      }
    };

    return comp;
  });
})(window.$flare);
</script>
