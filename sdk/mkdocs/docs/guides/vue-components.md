# Vue Components

Flare Hotspot uses [Vue.js](https://v2.vuejs.org) to build the user interface. But we are not using the standard build tools for Vue.js project since we need to support dynamic components from the plugins. Hence, the syntax for declaring vue components are slightly different. This guide will help you understand how to build and use Vue components in the Flare Hotspot project.

Take a look at the following example:

```html
<!-- resources/components/Welcome.vue -->

<template>
  <div>
    <h1>Welcome to Flare Hotspot</h1>
  </div>
</template>

<script>
define(function() {
    return {
        props: ['flareView'],
        template: template
    }
})
</script>

```
