<template>
    <div>
        <ul>
            <li v-for="(url, name) in flareView.data">
                <router-link :to="url">{{ name }}</router-link>
            </li>
        </ul>
    </div>
</template>

<script>
define(function () {
    return {
        props: ['flareView'],
        template: template,
        mounted: function () {
            console.log(this.$route);
        }
    };
});
</script>
