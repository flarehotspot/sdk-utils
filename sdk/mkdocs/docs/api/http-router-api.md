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

This method is used to generate the url for third-party plugin route name. This method accepts three arguments, the first argument is the plugin package name (e.g `com.flarego.core`), the second argument is the route name and the third argument is the route parameters (similar to [UrlForRoute](#urlforroute) method):

```go
url := api.Http().HttpRouter().UrlForPkgRoute("com.flarego.core", "portal.welcome", "name", "John")
```

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
}).Name("payments.options") // set the route name
```

### Post

This method is used to create a route for the `POST` http method. This method accepts tree (or more) arguments, the first argument is the route path, the second argument is the handler function and the third and subsequent arguments are the list of (optional) [middlewares](#middlewares). Take a look at the example below:

```go
router := api.Http().HttpRouter().AdminRouter()
router.Post("/settings/save", func(w http.ResponseWriter, r *http.Request) {
    // Handle the request
}).Name("settings.save") // set the route name
```

### Use

This method is used to add a [middleware](#middlewares) to the router. This method accepts a list of middlewares.
All routes defined after the `Use` method will use the middleware.

Below is using a middleware for plugin sub-router:
```go
middleware := func (next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // do something before the handler function
        next.ServeHTTP(w, r)
    })
}

api.Http().HttpRouter().PluginRouter().Group("/payments", func (subrouter sdkhttp.HttpRouterInstance) {
    subrouter.Use(middware)
})
```

Below is using a middleware for admin sub-router:
```go
middleware := func (next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // do something before the handler function
        next.ServeHTTP(w, r)
    })
}

api.Http().HttpRouter().AdminRouter().Group("/settings", func (subrouter sdkhttp.HttpRouterInstance) {
    subrouter.Use(middware)
})
```

In the examples above, the middleware is used to perform operations on the request before it reaches the handler function inside the sub-router.

## Middlewares

A middleware is a function of type `func(next http.Handler) http.Handler`. It is used to perform operations on the request before it reaches the handler function. Middlewares are functions that accept a http handler function and returns another http handler function.

### Declaring a middleware
Below is an example of a middleware:

```go
middleware := func (next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // do something before the handler function
        next.ServeHTTP(w, r)
    })
}
```

### Using a middleware
Then you can use the middleware in the route definition:

```go
router := api.Http().HttpRouter().AdminRouter()
router.Post("/settings/save", func(w http.ResponseWriter, r *http.Request) {
    // Handle the request
}, middleware) // use the middleware
```
