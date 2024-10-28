<template>
  <div>
    <router-view :data="data"></router-view>
  </div>
</template>

<script>
define(function () {
  return {
    props: ["data"],
    template: template
  };
});
</script>
