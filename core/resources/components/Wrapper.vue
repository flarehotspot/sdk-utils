<template>
    <wrapped-component :flare-view="view"></wrapped-component>
</template>

<script>
(function (window) {
    var $flare = window.$flare;
    var compPath = 'vue!<% .Data.HttpComponentFullPath %>';
    var dataPath = '<% .Data.HttpDataFullPath %>';

    define([compPath], function (comp) {
        return {
            template: template,
            components: {
                WrappedComponent: comp
            },
            data: function () {
                return { view: { loading: true, data: {}, errors: {} } };
            },
            mounted: function () {
                var self = this;
                var params = this.$route.params;
                var query = this.$route.query;
                var path = $flare.utils.vuePathToMuxPath(dataPath, params);
                path = $flare.utils.attachQueryParams(path, query);

                $flare.http
                    .get(path)
                    .then(function (data) {
                        self.view = { data: data, errors: {}, loading: false };
                    })
                    .catch(function (err) {
                        console.log(err);
                        self.view = { data: {}, errors: err, loading: false };
                    });
            }
        };
    });
})(window);
</script>
