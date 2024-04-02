<template lang="html">
    <div>

    <h1> Logs </h1>

    <div>{{ console.log(flareView.data) }}</div>
    <h1>{{ flareView.data}}</h1>
</div>
</template>

<script>
define(function () {
    return {
        template: template,
        props: ['flareView'],
        methods: {}
    };
});
</script>