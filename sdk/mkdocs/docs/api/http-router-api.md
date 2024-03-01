# HttpRouterApi

The `HttpRouterApi` is the backend for http routing in Flare Hotspot. The [VueRouterApi](./vue-router-api.md) uses the `HttpRouterApi` to generate the routes for the frontend. Each plugin are provided with a `HttpRouterApi` instance to generate their own routes.

## HttpRouterApi Methods

Below are the available methods in `HttpRouterApi`:

### PluginRouter

This method returns the [plugin router instance](#router-instance) for the plugin routes. Routes generated from the plugin router are accessible to all users. To get the pugin router instance, you can use the following code:

```go
router := api.Http().HttpRouter().PluginRouter()
```

### AdminRouter

This method returns the [admin router instance](#router-instance) for the admin routes. Routes generated from the admin router are prefixed with `/admin` and are only accessible to authenticated user [accounts](./accounts-api.md#account-instance). To get the admin router instance, you can use the following code:

```go
router := api.Http().HttpRouter().AdminRouter()
```

### UrlForRoute

This method is used to generate the url for the given plugin route name. This method accepts two arguments, the first argument is the route name and the second argument is a map of route parameters. The route parameters are key-value pairs. The example below generates a URL for the route name `portal.welcome` with a route path `/welcome/:name`:

```go
url := api.Http().HttpRouter().UrlForRoute("portal.welcome", "name", "John")
```

### UrlForPkgRoute

This method is used to generate the url for third-party plugin route name. This method accepts three arguments, the first argument is the plugin package name (e.g `com.flarego.core`), the second argument is the route name and the third argument is the route parameters (key-value pairs).

## Router Instance

Router instance is used to generate routes for the plugin. Below are the methods available in the router instance:

### Group

This method is used to create a group of routes with a common path prefix. This method accepts two arguments,
the first argument is the path prefix and the second argument is a function which accepts another router instance.
This can be used to nest more route groups. Take a look at the example below:

```go
// Get the router instance
router := api.Http().HttpRouter().PluginRouter()
router.Group("/payments", func (subrouter sdkhttp.HttpRouterInstance) {
    // Add route for /payments/received
    subrouter.Post("/receieved", func(w http.ResponseWriter, r *http.Request) {
        // Handle the request
    })
})
```

### Get

This method is used to create a route for the `GET` http method. This method accepts tree (or more) arguments, the first argument is the route path, the second argument is the handler function and the third and subsequent arguments are the list of (optional) [middlewares](#middlewares). Take a look at the example below:

```go
router := api.Http().HttpRouter().PluginRouter()
router.Get("/payments/options", func(w http.ResponseWriter, r *http.Request) {
    // Handle the request
})
```

### Post

## Middlewares
