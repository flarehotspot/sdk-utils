
# Basic Routing

## Portal Routes
Below is an example of how to define a [portal route](../api/vue-router.md#portalroute) using the [VueRouter.RegisterPortalRoutes](../api/vue-router.md#registerportalroutes) api method.

```go
package main

import (
	"net/http"

	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func main() {}

func Init(api sdkplugin.PluginApi) {
	// define the portal route
	portalRoute := sdkhttp.VuePortalRoute{
		RouteName: "portal.welcome",
		RoutePath: "/welcome/:name",
		Component: "portal/Welcome.vue",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			params := api.Http().MuxVars(r)
			name := params["name"]

			data := map[string]interface{}{
				"name": name,
			}

			api.Http().VueResponse().Json(w, data, 200)
		},
	}

    // register portal route
	api.Http().VueRouter().RegisterPortalRoutes(portalRoute)
}
```

## Admin Routes
Admin routes are very similar to portal routes, but they are used to display web pages in the admin panel. The difference between portal routes and admin routes is that admin routes cannot be accessed by unauthenticated users. To define an admin route, we use the [VueRouter.RegisterAdminRoutes](../api/vue-router.md#registeradminroutes) api method.

```go
// define admin route
adminRoute := sdkhttp.VueAdminRoute{
    RouteName: "admin.welcome",
    RoutePath: "/welcome/:name",
    Component: "admin/Welcome.vue",
    HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
        params := api.Http().MuxVars(r)
        name := params["name"]
        data := map[string]interface{}{
            "name": name,
        }
        api.Http().VueResponse().Json(w, data, 200)
    },
}
// register the admin route
api.Http().VueRouter().RegisterAdminRoutes(adminRoute)
```

Below is the brief definition of each fields used to define the [Portal Route](../api/vue-router.md#portalroute) and [Admin Route](../api/vue-router.md#adminroute).

## RouteName
This field can be used to reference this route in case we want to link this page from another page using the [HttpHelpers.VueRoutePath](../api/http-helpers.md#vueroutepath) method. Below is an example of creating a link from another page to the portal route we defined above.

```html
<!-- Create a link to the portal route named "portal.welcome" with a param "name" of value "Jhon" -->
<router-link :to='<% .Helpers.VueRoutePath "portal.welcome" "name" "Jhon" %>'>Go to welcome page</router-link>
```

## RoutePath
This field is used to match the URL in the browser which will trigger the portal route. Route params can be extracted using
[HttpApi.MuxVars](../api/http-api.md#muxvars) method.

```go
// handler func
func (w http.ResponseWriter, r *http.Request) {
    // get the route params
    params := api.Http().MuxVars(r)
    fmt.Println(params["name"]) // Jhon
}
```

## HandlerFunc
This field is used to define the handler function for the Vue.js component. The returned response from [VueResponse.Json](../api/vue-response.md#json) will be available in the Vue component in `flareView` [prop](https://v2.vuejs.org/v2/guide/components-props).

## Component
This field defines the location of the [Vue.js component](https://v2.vuejs.org/v2/guide/components) file to be displayed in the web page. Vue components are loaded from the `resources/components` directory under the root directory of your plugin. Below is an example of a Vue component that displays the json data from the [HandlerFunc](#handlerfunc).

```html
<!-- resources/components/portal/Welcome.vue -->

<template>
    <h1>Welcome {{ flareView.data.name }}</h1>
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

The `flareView` prop is automatically populated with the JSON data from the handler function defined in [HandlerFunc](#handlerfunc) field of the portal route. The `flareView` component prop has three fields, namely:

- `data`: The JSON data returned from the handler function.
- `loading`: A boolean value that indicates if the data is still loading.
- `error`: A string containing the error message if the data loading fails.

The `template` variable is a string containing the HTML code automatically extracted from the `<template>` tag. Note that there must be **ONLY ONE** root html tag of the template.
