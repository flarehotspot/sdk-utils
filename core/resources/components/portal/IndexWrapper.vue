<!--
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->

<template>
    <wrapped-component :flare-view="view" :portal-items="items"></wrapped-component>
</template>

<script>
(function (window) {
    var $flare = window.$flare;
    var compPath = 'vue!<% .Data.HttpComponentFullPath %>';
    var dataPath = '<% .Data.HttpDataFullPath %>';
    var portalItemsPath =
        '<% .Helpers.UrlForPkgRoute "com.flarego.core" "portal.items" %>';

    define([compPath], function (comp) {
        var reloadListener = null;

        return {
            template: template,
            components: {
                WrappedComponent: comp
            },
            data: function () {
                return { view: { loading: true, data: {}, error: null }, items: [] };
            },
            mounted: function () {
                var self = this;
                self.load();

                reloadListener = $flare.events.on(
                    'portal:items:reload',
                    function (items) {
                        self.items = items;
                    }
                );
            },
            beforeDestroy: function () {
                if (reloadListener) {
                    $flare.events.off('portal:items:reload', reloadListener);
                }
            },
            methods: {
                load: function () {
                    var self = this;
                    var params = this.$route.params;
                    var query = this.$route.query;
                    var path = $flare.utils.vuePathToMuxPath(dataPath, params);
                    path = $flare.utils.attachQueryParams(path, query);

                    var compData = $flare.http.get(path);
                    var portalItems = $flare.http.get(portalItemsPath);

                    Promise.all([compData, portalItems])
                        .then(function (values) {
                            var data = values[0];
                            var items = values[1];
                            self.view = {
                                data: data,
                                error: null,
                                loading: false
                            };
                            self.items = items;
                        })
                        .catch(function (err) {
                            console.error(err);
                            self.view = { data: {}, error: err, loading: false };
                        });
                }
            }
        };
    });
})(window);
</script>
