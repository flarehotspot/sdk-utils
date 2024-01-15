<template>
  <div>
    <h1>{{ message }}: {{ counter }}</h1>
    <button @click="counter++">Click me</button>
    <div v-for="index in counter" :key="index">
      <test-component></test-component>
    </div>
  </div>
</template>

<script>
define(["com.flarego.core/captive-portal/helpers"], function (h) {
  return {
    components: {
      "test-component": h.vueLoader("/components/Test")
    },
    data: function () {
      return {
        counter: 0,
        message: "Hello World!"
      };
    },
    template: template
  };
});
</script>
