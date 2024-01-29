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
                var path = substitutePathParams(dataPath, params);

                $flare.http
                    .get(path)
                    .then(function (data) {
                        self.view = { data: data, errors: {}, loading: false };
                    })
                    .catch(function (err) {
                        console.log(err)
                        self.view = { data: {}, errors: err, loading: false };
                    });
            }
        };
    });

    function substitutePathParams(path, params) {
        // Regular expression to match {param} in the path
        const paramRegex = /\{([^}]+)\}/g;

        // Replace each {param} with the corresponding value from the params object
        const substitutedPath = path.replace(paramRegex, function (_, paramName) {
            // If the param exists in the params object, use its value, otherwise, keep the original {param}
            return params.hasOwnProperty(paramName)
                ? params[paramName]
                : '{' + paramName + '}';
        });

        return substitutedPath;
    }
})(window);
</script>
