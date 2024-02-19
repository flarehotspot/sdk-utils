+++
title = "Http API"
description = "Http API provides access to the system's HTTP server."
date = 2021-05-01T08:00:00+00:00
updated = 2021-05-01T08:00:00+00:00
draft = false
weight = 20
sort_by = "weight"
template = "docs/page.html"

[extra]
lead = "Http API provides access to the system's HTTP server."
toc = true
top = false
+++

# Overview

The `HttpApi` is used to access various HTTP server functionalities including authentication, routing, and http responses.

# Methods
First, get an instance of the `HttpApi` from the [PluginApi](../plugin-api):
```go
package main
// imports...
func Init(api sdkplugin.PluginApi) {
    httpApi := api.Http()
}
```

## Auth
It returns an instance of the [HttpAuth](../http-auth).
```go
auth := httpApi.Auth()
```

## Helpers
It returns an instance of the [HttpHelpers](../http-helpers).
```go
helpers := httpApi.Helpers()
```

## HttpRouter
It returns an instance of [HttpRouter](../http-router).
```go
httpRouter := httpApi.Router()
```

## VueRouter
It returns an instance of [VueRouter](../vue-router).
```go
vueRouter := httpApi.VueRouter()
```

## Middlewares
Returns an instance of [Middlewares](../middlewares).
```go
middlewares := httpApi.Middlewares()
```

## HttpResponse
Returns an instance of [HttpResponse](../http-response).
```go
httpResponse := httpApi.HttpResponse()
```

## VueResponse
Returns an instance of [VueResponse](../vue-response).
```go
vueResponse := httpApi.VueResponse()
```

## GetDevice
Get the device information from the http request. It returns and instance of [ClientDevice](../client-device) and an `error`.
This must be used in conjuction with the [Device middleware](../middlewares/#device).
```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    // other logic...
    device, err := httpApi.GetDevice(r)
    if err != nil {
        // handle error
    }
    fmt.Println(device) // ClientDevice
}
```

## MuxVars
Returns a map (`map[string]string`) of mux variables from the request path. For example, if the route pattern is `/sessions/{id}` and the request path is `/sessions/1`, get the `id` param with:
```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    // other logic...
    vars := httpApi.MuxVars(r) // map[string]string
    id := vars["id"]
    fmt.Println(id) // "1"
}
```

## GetAdminNavs
Returns a slice of [AdminNavList](#admin-nav-list)
```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    // other logic...
    navList := httpApi.GetAdminNavs(r)
    fmt.Println(navList) // []AdminNavList
}
```
## GetPortalItems
Returns a slice of [PortalItem](#portal-item)
```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    // other logic...
    portalItems := httpApi.GetPortalItems(r)
    fmt.Println(portalItems) // []PortalItem
}
```

# AdminNavList {#admin-nav-list}
`AdminNavList` is a list of items for the admin navigation. It has the following fields:
```go
type AdminNavList struct {
	Label string         `json:"label"`
	Items []AdminNavItem `json:"items"`
}
```

# AdminNavItem {#admin-nav-item}
`AdminNavItem` is an item for the admin navigation. It has the following fields:
```go
type AdminNavItem struct {
	Category     INavCategory `json:"category"`
	Label        string       `json:"label"`
	VueRouteName string       `json:"route_name"`
	VueRoutePath string       `json:"route_path"`
}
```

# PortalItem {#portal-item}
`PortalItem` is an item for the portal navigation. It has the following fields:
```go
type PortalItem struct {
	IconUri      string `json:"icon_uri"`
	Label        string `json:"label"`
	VueRouteName string `json:"route_name"`
	VueRoutePath string `json:"route_path"`
}
```
