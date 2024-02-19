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
(function () {
    define(function () {
        return {
            props: ['flareView'],
            template: template
        };
    });
})();
</script>
