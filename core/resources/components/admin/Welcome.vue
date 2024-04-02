<template>
    <div>
        <h1>Welcome {{ flareView.data.name }}. You are an admin.</h1>
    </div>
</template>
<script>
    define(function(){
        // return the vue component definition
        return {
            props: ['flareView'],
            template: template,
            mounted: function(){
                console.log(this.flareView)
                // { data: { name: "John" }, error: null, loading: false }
            }
        }
    })
</script>