+++
title = "Http Router"
description = "Http Router allows plugins to add routing to the system's HTTP server."
date = 2021-05-01T08:00:00+00:00
updated = 2021-05-01T08:00:00+00:00
draft = false
weight = 20
sort_by = "weight"
template = "docs/page.html"

[extra]
lead = "Http Router allows plugins to add routing to the system's HTTP server."
toc = true
top = false
+++

# HttpRouter

## Overview

The `HttpRouter` is used to define HTTP routes, middlewares and their handlers for generic http requests within your plugin.
It is different from the [VueRouter](../vue-router). The later is used to define routes for the vue application, the frontend framework used in our system.

# Methods

First, you need to get an instance of `HttpRouter` from [PluginApi]('../plugin-api'):
```go
package main
// imports...
func Init(api sdkplugin.PluginApi) {
    router := api.HttpRouter()
}
```

## AdminRouter
Returns a [Router](#router-instance) instance. This router have an [AdminAuth](../middlewares/#admin-auth) middleware by default to make sure HTTP requests are coming from an authenticated admin user. Route paths registered using the admin router have `/admin` prefix.

```go
adminRouter := router.AdminRouter()
```

## PluginRouter
Returns a [Router](#router-instance) instance. This router doesn't have any middleware by default.
```go
pluginRouter := router.PluginRouter()
```

## UrlForRoute
It returns the URL to the route with the given [PluginRouteName](#plugin-route-name) and the given param pairs.
```go
pluginRoute := sdkhttp.PluginRouteName("my-plugin-route")
url := router.UrlForRoute(pluginRoute, "param1", "value1", "param2", "value2")
```

## MuxRouteName
It accepts a [PluginRouteName](#plugin-route-name) and returns a [MuxRouteName](#mux-route-name)
```go
muxRouteName := router.MuxRouteName(sdkhttp.PluginRouteName("my-route-name"))
```

## UrlForMuxRoute
It returns the URL to the route with the given [MuxRouteName](#mux-route-name) and the given param pairs.
```go
pluginRoute := sdkhttp.RouteName("my-plugin-route")
muxRouteName := router.MuxRouteName(pluginRoute)
url := router.UrlForMuxRoute(muxRouteName, "param1", "value1", "param2", "value2")
```

---

# Router Instance {#router-instance}
A `RouterInstance` is returned from [router.AdminRouter](#adminrouter) or [router.PluginRouter](#pluginrouter). This is where you define your routes, middlewares and handlers. Only `GET` and `POST` requests are supported.

## Get
Define a `GET` route for your plugin. Below is an example of a `GET` route with a middleware and a route name. It returns an [HttpRoute](#http-route) instance.
```go
myHandler := func (w http.ResponseWriter, r *http.Request) {
    // your handler logic
}
mydMiddleware := func (next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // your middleware logic
        next.ServeHTTP(w, r)
    })
}
routeName := sdkhttp.PluginRouteName("my-route")
rtr := router.PluginRouter() // or can also be router.AdminRouter()
rtr.Get("/my-route-path", myHandler, myMiddleware).Name(routeName)
```

## Post
Define a `POST` route for your plugin. Below is an example of a `POST` route with a middleware and a route name. It returns [HttpRoute](#http-route) instance.
```go
myHandler := func (w http.ResponseWriter, r *http.Request) {
    // your handler logic
}
myMiddleware := func (next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // your middleware logic
        next.ServeHTTP(w, r)
    })
}
routeName := sdkhttp.PluginRouteName("my-route")
rtr := router.PluginRouter() // or can also be router.AdminRouter()
rtr.Post("/my-route-path", myHandler, myMiddleware).Name(routeName)
```
---

# Http Route
An `HttpRoute` is returned from [RouterInstance.Get](#get) or [RouterInstance.Post](#post) methods. It is used to define a route name.

## Name
Define a [PluginRouteName](#plugin-route-name) for the route.
```go
route := rtr.Get("/my-route-path", myHandler)
route.Name(sdkhttp.PluginRouteName("my-route-name"))
```

---

# Route Names

## PluginRouteName {#plugin-route-name}
A `PluginRouteName` is a route name defined in the routes within your plugin. For example, if you defined a route:
```go
    routeName := sdkhttp.PluginRouteName("my-route-name")
    rtr := router.PluginRouter()
    rtr.Get("/my-route-path", myHandler).Name(routeName)
```
Then, "my-route-name" is the `PluginRouteName`.

## MuxRouteName {#mux-route-name}
A `MuxRouteName` are route names which can be used to reference a page outside your plugin. Internally, a [PluginRouteName](#plugin-route-name) is converted into a `MuxRouteName` prefixed with [plugin package name](../plugin-api/#pkg). Internal routes in the system are all using `MuxRouteName`. Below are the available `MuxRouteName` in the system:

TODO: Add available MuxRouteNames

