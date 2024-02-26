<!--
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->

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
      "test-component": h.VueLoader("/components/Test")
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
