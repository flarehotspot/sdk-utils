# IHttpMiddlewares

## About Middlewares {#about-middlewares}

Middlewares are used to perform operations before the handler function is executed. These are functions that accept `http.Handler` and return `http.Handler`. Below is an example of how to define a middleware:
```go title="main.go"
middleware := func (next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // do something before the handler function
        next.ServeHTTP(w, r)
    })
}

handler := func (w http.ResponseWriter, r *http.Request) {
    // handle http request
}
```

Then you can use the middleware in the http route:
```go title="main.go"
router := api.Http().HttpRouter().PluginRouter()
router.Get("/welcome", handler, middleware).Name("welcome")
```

## Built-in Middlewares {#built-in-middlewares}

To get an instance of the built-in `IHttpMiddlewares`:

```go
httpMw := api.Http().Middlewares()
fmt.Println(httpMw) // IHttpMiddlewares
```

Below are the list of available built-in middlewares in `IHttpMiddlewares`:

### AdminAuth {#admin-auth}

It returns a middleware that ensures that only authenticated admins can access the route.

```go
pluginRouter := api.Http().HttpRouter().PluginRouter()
authMw := httpMw.AdminAuth()
handler := func(w http.ResponseWriter, r *http.Request) {
    // handle the http request...
}
pluginRouter.Get("/protected-page", handler, authMw)
```

### CacheResponse {#cache-response}

It returns a middleware that caches the response throughout the duration of the application lifetime which can improve system performance.

```go
pluginRouter := api.Http().HttpRouter().PluginRouter()
cacheMw := httpMw.CacheResponse()
handler := func(w http.ResponseWriter, r *http.Request) {
    // handle the http request...
}
pluginRouter.Get("/some-generated-image.png", handler, cacheMw)
```

### PendingPurchase

It returns a middleware that redirects the user to the pending order payment page when a pending purchase request is present.


```go
pluginRouter := api.Http().HttpRouter().PluginRouter()
pendingMw := httpMw.PendingPurchase()
handler := func(w http.ResponseWriter, r *http.Request) {
    // handle the http request...
}
pluginRouter.Get("/some-checkout-page", handler, pendingMw)
```
