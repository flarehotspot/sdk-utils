<template lang="html">
  <div>
    <h1>Admin Themes:</h1>
    <div v-for="list in adminThemes" :key="list.theme_name">
      <input type="radio" :id="list.theme_pkg" :value="list.theme_pkg" v-model="selectedAdminTheme" />
      <label :for="list.theme_pkg">{{ list.theme_name }}</label>
    </div>

    <h1>Portal Themes:</h1>
    <div v-for="list in portalThemes" :key="list.theme_name">
      <input type="radio" :id="list.theme_pkg" :value="list.theme_pkg" v-model="selectedPortalTheme" />
      <label :for="list.theme_pkg">{{ list.theme_name }}</label>
    </div>

    <button type="button" @click="changeTheme">Save Changes</button>
  </div>
</template>
<script>
define( function() {
    return{
      data(){
        return{
          selectedAdminTheme:null,
          selectedPortalTheme:null
        };
      },
      template: template,
      props:['flareView'],
      computed: {
        adminThemes: function() {
          return this.flareView.data.theme_admin;
        },
        portalThemes: function() {
          return this.flareView.data.theme_portal;
        }
      },
      methods: {
        changeTheme: function(){
          var savedData={
            admin: this.selectedAdminTheme,
            portal:this.selectedPortalTheme
          }

          window.$flare.http.post('<% .Helpers.UrlForRoute "admin.theme.save" %>', savedData).catch(function (err) {
            console.log(err)
          })
        }
      },
    }
})
</script>