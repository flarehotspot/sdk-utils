# HttpMiddlewares

## 1. Middlewares

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

## 2. Built-in Middlewares

There are some built-in middlewares that you can use:

| Middleware | Description |
|------------|-------------|
| AdminAuth | This middleware ensures that only authenticated admins can access the route. |
| CacheResponse | This middleware caches the response throughout the duration of the application runtime.
| PendingPurchase | This middleware redirects the user to the pending order payment page, if any.

For example, to use the `AdminAuth` middleware:

```go
adminAuthMw := api.Http().Middlewares().AdminAuth()
router.Get("/welcome", handler, adminAuthMw).Name("welcome")
```

