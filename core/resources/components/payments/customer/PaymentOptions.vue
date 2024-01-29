<template>
    <div>
        <ul>
            <li v-for="(url, name) in flareView.data">
                <router-link :to="attachTokenToPath(url)">{{ name }}</router-link>
            </li>
        </ul>
    </div>
</template>

<script>
(function ($flare) {
    define(function () {
        return {
            props: ['flareView'],
            template: template,
            mounted: function () {
                console.log(this.$route);
            },
            methods: {
                attachTokenToPath: function (path) {
                    var self = this;
                    var token = self.$route.query.token;
                    return $flare.utils.attachQueryParams(path, { token: token });
                }
            }
        };
    });
})(window.$flare);
</script>
