# Vue Components

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

## The `flareView` prop

The `flareView` prop is automatically populated with the JSON data from the handler function defined in [HandlerFunc](#routes-and-links.md#handlerfunc) field of the portal/admin route. The `flareView` component prop has three fields, namely:

- `data`: The JSON data returned from [VueResponse.Json](../api/vue-response.md#json) method called inside the [handler function](./routes-and-links.md#handlerfunc).
- `loading`: A boolean value that indicates if the data is still loading.
- `error`: A string containing the error message if the data loading fails.

## The `template` variable

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

## Template helpers
Aside from the [HttpHelpers.VueRoutePath](../api/http-helpers.md#vueroutepath) method we used to create a link, there are other useful methods within the [HttpHelpers](../api/http-helpers.md) API. The [HttpHelpers](../api/http-helpers.md) can be accessed anywhere inside the component as `.Helpers` (notice the dot prefix) enclosed by `<%` and `%>` delimiters. Visit the [HttpHelpers](../api/http-helpers.md) API documentation to learn more.

For example, to build a link to another route, you can use the `HttpHelpers.VueRoutePath` method as shown below:
```html
<router-link :to='<% .Helpers.VueRoutePath "portal.welcome" %>'>Welcome</router-link>
```

## Browser Compatibility
Since we are not using standard build tools like webpack or vite, it is recommended to use basic form of javascript and css to ensure compatibility with older browsers. For example, use `var` instead of `let` or `const` and use `function` instead of arrow functions.
