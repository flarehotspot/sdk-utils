+++
title = "Creating a portal page"
description = "In this tutorial, we will create our first page to display to the user."
date = 2021-05-01T08:00:00+00:00
updated = 2021-05-01T08:00:00+00:00
draft = false
weight = 110
sort_by = "weight"
template = "docs/page.html"

+++

# Creating Captive Portal Page

In this tutorial, we will create our first page to display to the user.
The [wifi captive portal](https://en.wikipedia.org/wiki/Captive_portal) pages (and [admin pages](../creating-admin-page)) are defined using the [VueRouter](../api/vue-router/) API.
A route component is a [Vue.js(v2) component](https://v2.vuejs.org/v2/guide/components) with the corresponding route name and path.

Below is a simple example of a captive portal page that displays a welcome message to the user.

```go
// main.go
package main

import (
    "net/http"
    sdkplugin "github.com/flarehotspot/core/sdk/api/plugin"
    sdkhttp "github.com/flarehotspot/core/sdk/api/http"
)

func Init(api sdkplugin.PluginApi) {
	api.Http().VueRouter().RegisterPortalRoutes(sdkhttp.VuePortalRoute{
		RouteName:   "portal.welcome",
		RoutePath:   "/welcome/:name",
		Component:   "portal/Welcome.vue",
		HandlerFunc: func (w http.ResponseWriter, r *http.Request) {
		    name := api.Http().MuxVars(r)["name"]
            res := api.Http().VueResponse()
            data := map[string]string{
                "name": name,
            }
            res.Json(w, data, 200)
        },
	})
}
```

Now let's create the `Welcome.vue` component file in the `resources/components/portal` directory of your plugin.

```html
<!-- resources/components/portal/Welcome.vue -->
<template>
    <div>
        <h1>Welcome {{ flareView.data.name }}</h1>
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
```

In this example, we are registering a captive portal route with the name `portal.welcome` and the path `/welcome/:name`.
The `Component` field specifies the [Vue component](https://v2.vuejs.org/v2/guide/components) to render when the user navigates to this route.

The `HandlerFunc` field is a function that is called when the user navigates to the route.
The returned data from [VueResponse.Json](../api/vue-response/#json) is automatically passed to the vue component as `flareView` component prop.

Route params can be defined using a colon (`:`) prefix. In this example, we defined a route param called `:name` which is used to display the welcome message in our vue component.

The `flareView` component prop has three fields namely:

- `data` - the json data return from the handler function when using [VueResponse.Json](../api/vue-response/#json)
- `error` - an error that occurred when fetching the data from the handler
- `loading` - a boolean that indicates if the data is still being fetched from the handler

The `template` field of the component is assigned with the `template` variable. The template variable contains the html string inside the `<template>` html tag.

Now that we've registered our first page to the router, we should rebuild and run the SDK:
```bash
docker compose restart app
```

The generated route paths are prefixed with the plugin `package` and `version` fields from [plugin.json](../api/plugin-json) file.
So to visit the page, you can navigate to the following URL in your browser:

[http://localhost:3000/#/com.sample.plugin/0.0.1/welcome/John](http://localhost:3000/#com.sample.plugin/0.0.1/welcome/John)

