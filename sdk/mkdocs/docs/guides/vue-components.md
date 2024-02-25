# Vue Components

## Introduction

Flare Hotspot uses [Vue.js](https://v2.vuejs.org) to build the user interface. But we are not using the standard build tools for Vue.js project since we need to support dynamic components from the plugins. Hence, the syntax for declaring vue components are slightly different. This guide will help you understand how to build and use Vue components in the Flare Hotspot project. Vue components are placed in the `resources/components` directory in your plugin.

Take a look at the following example:

```html title="resources/components/portal/Welcome.vue"
<template>
    <div>
        <h1>Welcome {{ flareView.data.name }}</h1>
        <p>Some other text</p>
    </div>
</template>

<script>
define(function () {
    return {
        props: ['flareView'],
        template: template,
    }
})
</script>
```

### The `flareView` prop

The `flareView` prop is automatically populated with the JSON data from the handler function defined in [HandlerFunc](#handlerfunc) field of the portal/admin route. The `flareView` component prop has three fields, namely:

- `data`: The JSON data returned from [VueResponse.Json](../api/vue-response.md#json) method called inside the [handler function](./basic-routing.md#handlerfunc).
- `loading`: A boolean value that indicates if the data is still loading.
- `error`: A string containing the error message if the data loading fails.

### The `template` variable

The `template` variable is a string containing the HTML code automatically extracted from the `<template>` tag.

!!! warning "Important"
    Note that there must be **only one root html tag** of the template. The following template will not work:
    ```html
    <template>
        <h1>Welcome {{ flareView.data.name }}</h1>
        <p>Some other text</p> <!-- the <p> tag will not be displayed -->
    </template>
    ```

    Below is the correct way:
    ```html
    <template>
        <div>
            <h1>Welcome {{ flareView.data.name }}</h1>
            <p>Some other text</p>
        </div>
    </template>
    ```

## Go html/template package

### Helpers
Aside from the [Http.VueRoutePath](../api/http-helpers.md#vuerouetpath) method we used to create a link, there are other useful methods within the [HttpHelpers](../api/http-helpers.md) package. The `HttpHelpers` can be accessed in the views as `.Helpers` (notice the dot prefix).

You can check the [HttpHelpers](../api/http-helpers.md) documentation for more information.
