<template lang="html">
    <div>
        <h1>
            <% .Helpers.Translate "label" "admin_themes" %>:
        </h1>
        <div v-for="theme in adminThemes" :key="theme.pkg">
            <input type="radio" :value="theme.pkg" v-model="flareView.data.themes_config.admin" />
            <label :for="theme.pkg">{{ theme.name }}</label>
        </div>

        <h1>
            <% .Helpers.Translate "label" "portal_themes" %>:
        </h1>
        <div v-for="theme in portalThemes" :key="theme.pkg">
            <input type="radio" :value="theme.pkg" v-model="flareView.data.themes_config.portal" />
            <label :for="theme.pkg">{{ theme.name }}</label>
        </div>

        <button type="button" @click="changeTheme">Save Changes</button>
    </div>
</template>

<script>
define(function () {
    return {
        template: template,
        data: function () {
            return {
                adminThemes: [],
                portalThemes: [],
                config: {
                    admin: '',
                    portal: ''
                }
            };
        },
        mounted: function () {
            var self = this;
            $flare.http
                .get('<% .Helpers.UrlForRoute "admin.themes.index" %>')
                .then(function (data) {
                    console.log(data)
                    self.adminThemes = data.admin_themes;
                    self.portalThemes = data.portal_themes;
                    self.config = data.themes_config;
                });
        },
        methods: {
            changeTheme: function () {
                var data = this.flareView.data;
                var savedData = {
                    admin: data.themes_config.admin,
                    portal: data.themes_config.portal
                };

                window.$flare.http
                    .post('<% .Helpers.UrlForRoute "admin.themes.save" %>', savedData)
                    .then(function () {
                        $flare.notify.success('Themes saved successfully.');
                    })
                    .catch(function (err) {
                        console.error(err);
                        $flare.notify.error(err);
                    });
            }
        }
    };
});
</script>
