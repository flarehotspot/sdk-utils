<!--
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->

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
                return { view: { loading: true, data: {}, error: null } };
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
                        self.view = { data: data, error: null, loading: false };
                    })
                    .catch(function (err) {
                        console.log(err);
                        self.view = { data: {}, error: err, loading: false };
                    });
            }
        };
    });
})(window);
</script>
