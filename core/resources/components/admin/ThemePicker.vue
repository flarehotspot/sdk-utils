<template lang="html">
    <div>
        <h1><% .Helpers.Translate "label" "admin_themes" %>:</h1>
        <div v-for="theme in flareView.data.admin_themes" :key="theme.pkg">
            <input type="radio" :value="theme.pkg" v-model="flareView.data.themes_config.admin" />
            <label :for="theme.pkg">{{ theme.name }}</label>
        </div>

        <h1><% .Helpers.Translate "label" "portal_themes" %>:</h1>
        <div v-for="theme in flareView.data.portal_themes" :key="theme.pkg">
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
        props: ['flareView'],
        data: function(){
            return {
                sample_data: 'sample data'
            };
        },
        mounted: function(){
            console.log('mounted');
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
