+++
title = "Vue Router"
description = "Http API provides access to the system's HTTP server."
date = 2021-05-01T08:00:00+00:00
updated = 2021-05-01T08:00:00+00:00
draft = false
weight = 100
sort_by = "weight"
template = "docs/page.html"

[extra]
lead = "Http API provides access to the system's HTTP server."
toc = true
top = false
+++

# Overview

The `HttpApi` is used to access various HTTP server functionalities including authentication, routing, and http responses.

# Http API
First, get an instance of the `HttpApi` from the [PluginApi]('../plugin-api'):
```go
package main
// imports...
func Init(api sdkplugin.PluginApi) {
    httpApi := api.Http()
}
```

## Auth
It returns an instance of the [HttpAuth](../auth-api).
```go
package main
// imports...
func Init(api sdkplugin.PluginApi) {
    httpApi := api.Http()
    authApi := httpApi.Auth()
    fmt.Println(authApi) // HttpAuth
}
```

## HttpRouter {#http-router-method}
It returns an instance of [HttpRouter](../http-router).
```go
func Init(api sdkplugin.PluginApi) {
    httpApi := api.Http()
    router := httpApi.HttpRouter()
    fmt.Println(router) // HttpRouter
}
```

---

# HttpRouter
