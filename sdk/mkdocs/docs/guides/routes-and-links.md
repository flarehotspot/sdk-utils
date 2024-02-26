# Routes and Links
Routes are used to handle user navigation by matching the requested URL to a [RoutePath](#routepath). A [link](#creating-a-link) is a form of clickable element in the web page that redirects a user to a certain URL and eventually triggering the matched route.

## Registering Routes

### Portal Routes {#portal-routes}
Below is an example of how to register a [portal route](../api/vue-router.md#portalroute) using the [VueRouter.RegisterPortalRoutes](../api/vue-router.md#registerportalroutes) api method.

```go title="main.go"
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
        Middlewares: []func(http.Handler) http.Handler{},
	}
	// register portal route
	api.Http().VueRouter().RegisterPortalRoutes(portalRoute)
}
```

### Admin Routes {#admin-routes}
Admin routes are very similar to portal routes, but are only accessible by authenticated user accounts. To define an admin route, we use the [VueRouter.RegisterAdminRoutes](../api/vue-router.md#registeradminroutes) api method.

```go title="main.go"
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
    Middlewares: []func(http.Handler) http.Handler{},
    PermitFn: func(perms []string) bool {
        // check if the user has the required permissions
        return true
    },
}
// register the admin route
api.Http().VueRouter().RegisterAdminRoutes(adminRoute)
```

## Route Definition {#route-definition}

Below is the brief definition of each fields used to define the [Portal Route](../api/vue-router.md#portalroute) and [Admin Route](../api/vue-router.md#adminroute).

### RouteName (required) {#routename}
This field can be used to reference this route in case we want to link this page from another page using the [HttpHelpers.VueRoutePath](../api/http-helpers.md#vueroutepath) method. Learn more about [creating a link](./creating-a-link.md).

### RoutePath (required) {#routepath}
This field is used to match the URL in the browser which will trigger the portal route. Route params can be extracted using
[HttpApi.MuxVars](../api/http-api.md#muxvars) method. For example, to get the `name` param from the route path `/welcome/:name`, you would do:

```go title="main.go"
// handler func
func (w http.ResponseWriter, r *http.Request) {
    // get the route params
    params := api.Http().MuxVars(r)
    fmt.Println(params["name"]) // Jhon
}
```

### Component (required) {#component}
This field defines the location of the [Vue Component](./vue-components.md) file to be displayed in the web page. Vue components are loaded from the `resources/components` directory of your plugin. Learn more about [Vue Components](./vue-components.md).

### HandlerFunc (optional) {#handlerfunc}
This field is used to define the handler function for the [Vue Component](./vue-components.md). The returned response from [VueResponse.Json](../api/vue-response.md#json) will be available in the Vue component in `flareView` [prop](https://v2.vuejs.org/v2/guide/components-props). A handler function is a function that accepts `http.ResponseWriter` and `*http.Request` arguments:

```go title="main.go"
func (w http.ResponseWriter, r *http.Request) {
    // send data to the vue component
    api.Http().VueResponse().Json(w, map[string]interface{}{"name": "Jhon"}, 200)
}
```

!!! note
    If the handler function is not defined, the server will just return `null` json data.

### Middlewares (optional) {#middlewares}
Middlewares are used to perform operations before the handler function is executed. Middlewares are functions that accept `http.Handler` and return `http.Handler`. Below is an example of how to define a middleware:
```go title="main.go"
mw := func (next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // do something before the handler function
        next.ServeHTTP(w, r)
    })
}
```

Then you can use the middleware in the route definition:
```go title="main.go"
portalRoute := sdkhttp.VuePortalRoute{
    // other fields...
    Middlewares: []func(http.Handler) http.Handler{mw},
}
```

### PermitFn (optional) {#permitfn}
This field is applicable only to admin routes. This function is used to validate the to access the admin route. The function accepts a slice of strings which contains the permissions of the account that's accessing the page. It's up to you to validate if the user can access the page. The function should return `true` if the user has the required permissions, otherwise `false`.
```go title="main.go"
permit := func (perms []string) bool {
    // check if the user has the required permissions
    return true
}
```

Then you can use the permit function in the route definition:
```go title="main.go"
adminRoute := sdkhttp.VueAdminRoute{
    // other fields...
    PermitFn: permit,
}
```

## Creating a Link {#creating-a-link}

### RouterLink {#routerlink}
A router link is a vue component that's part of the official [vue-router](https://github.com/vuejs/vue-router) package. We can create a link to a [portal route](./routes-and-links.md#portal-routes) or an [admin route](./routes-and-links.md#admin-routes) by using the [HttpHelpers.VueRoutePath](../api/http-helpers.md#vueroutepath) method.

```html title="AnotherPage.vue"
<router-link :to='<% .Helpers.VueRoutePath "portal.welcome" "name" "Jhon" %>'>Go to welcome page</router-link>
```
This creates a link to the portal route named `portal.welcome` with a param `name` of value `Jhon`.

### Route Params {#route-params}
Route params can be passed to the [Helpers.VueRoutePath](../api/http-helpers.md#vueroutepath) as key-value pairs. For example, if you have a route path `/users/:user_id/posts/:post_id` and the [name](./routes-and-links.md#routename) of the route is `user.posts`, this is how you can create a link to that route with params:
```html
<router-link :to='<% .Helpers.VueRoutePath "user.posts" "user_id" "1" "post_id" "2" %>'>
    User posts
</router-link>
```

Route params can be retreived in the handler function using the [HttpApi.MuxVars](../api/http-api.md#muxvars) method:
```go title="main.go"
func (w http.ResponseWriter, r *http.Request) {
    params := api.Http().MuxVars(r)
    fmt.Println(params["user_id"]) // 1
    fmt.Println(params["post_id"]) // 2
}
```
