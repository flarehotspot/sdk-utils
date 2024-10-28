# Routes and Links
Routes are used to handle user navigation by matching the requested URL from the browser to a [RoutePath](#routepath) in your application. A [link](#creating-a-link) is a form of clickable element in the web page that redirects a user to a certain URL and eventually triggering the matched route.

## 1. Registering Routes {#registering-routes}
The vue routes are divided into two types: [portal routes](#portal-routes) and [admin routes](#admin-routes). Portal routes are accessible to all users, while admin routes are only accessible to authenticated user accounts.

### Portal Routes {#portal-routes}
Below is an example of how to register a [portal route](../api/vue-router-api.md#portalroute) using the [VueRouterApi.RegisterPortalRoutes](../api/vue-router-api.md#registerportalroutes) api method.

```go title="main.go"
package main

import (
	"net/http"

	sdkhttp "sdk/api/http"
	sdkplugin "sdk/api/plugin"
)

func main() {}

func Init(api sdkplugin.PluginApi) {
	// define the portal route
	portalRoute := sdkhttp.VuePortalRoute{
		RouteName: "portal:welcome",
		RoutePath: "/welcome/:name",
		Component: "portal/Welcome.vue",
        Middlewares: []func(http.Handler) http.Handler{},
	}
	// register portal route
	api.Http().VueRouter().RegisterPortalRoutes(portalRoute)
}
```

See [Route Fields](#route-fields) for the definition of each fields used to define a portal route.

### Admin Routes {#admin-routes}
Admin routes are very similar to [portal routes](#portal-routes), but are only accessible by authenticated user accounts. To define an admin route, we use the [VueRouterApi.RegisterAdminRoutes](../api/vue-router-api.md#registeradminroutes) api method.

```go title="main.go"
// define admin route
adminRoute := sdkhttp.VueAdminRoute{
    RouteName: "admin:welcome",
    RoutePath: "/welcome/:name",
    Component: "admin/Welcome.vue",
    PermitFn: func(perms []string) bool {
        // check if the user has the required permissions
        return true
    },
}
// register the admin route
api.Http().VueRouter().RegisterAdminRoutes(adminRoute)
```

See [Route Fields](#route-fields) for the definition of each fields used to define an admin route.

## 2. Route Fields {#route-fields}

Below is the brief definition of each fields used to define the [Portal Route](../api/vue-router-api.md#portalroute) and [Admin Route](../api/vue-router-api.md#adminroute).

### RouteName (required) {#routename}
This field can be used to reference this route in case we want to link this page from another page using the [HttpHelpers.VueRoutePath](../api/http-helpers.md#vueroutepath) method. Learn more about [creating a link](#creating-a-link).

### RoutePath (required) {#routepath}
This field is used to match the URL in the browser to [portal](#portal-routes) or [admin](#admin-routes) route. Route params can be extracted using
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
This field defines the location of the [Vue Component](./vue-components.md) file to be displayed in the web page. Vue components are loaded from the `resources/components` directory of your plugin.

### PermitFn (optional) {#permitfn}
This field is applicable only to admin routes. This function is used to validate access to the admin route. The function accepts a slice of strings which contains the [permissions](../api/accounts-api.md#permissions-sec) of the account that's currently trying to access the page. It's up to you to validate if the user can access the page. The function should return `true` if the user has the required permissions, otherwise `false`.
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

## 3. Creating a Link {#creating-a-link}

### RouterLink {#routerlink}
A router link is a vue component that's part of the official [vue-router](https://github.com/vuejs/vue-router) package. We can create a link to a [portal route](./routes-and-links.md#portal-routes) or an [admin route](./routes-and-links.md#admin-routes) by using the [HttpHelpers.VueRoutePath](../api/http-helpers.md#vueroutepath) method.

```html title="AnotherPage.vue"
<router-link to='<% .Helpers.VueRoutePath "portal:welcome" "name" "Jhon" %>'>Go to welcome page</router-link>
```
This creates a link to the portal route named `portal:welcome` with a param `name` of value `Jhon`.

### Route Params {#route-params}
Route params can be passed to the [Helpers.VueRoutePath](../api/http-helpers.md#vueroutepath) as key-value pairs. For example, if you have a route path `/users/:user_id/posts/:post_id` and the [name](./routes-and-links.md#routename) of the route is `user.posts`, this is how you can create a link to that route with params:
```html
<router-link to='<% .Helpers.VueRoutePath "user.posts" "user_id" "1" "post_id" "2" %>'>
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
