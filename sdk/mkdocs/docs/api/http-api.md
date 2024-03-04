# HttpApi

The `HttpApi` is used to access various HTTP server functionalities including authentication, routing, and http responses.

## HttpApi Methods

The following are the available methods in `HttpApi`.

### GetClientDevice

Get the [client device](./client-device.md) info from the http request:

```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    device := api.Http().Helpers().GetClientDevice(r)
    fmt.Println(device) // ClientDevice
}
```

### Auth

It returns an instance of the [HttpAuth](./http-auth.md).

```go
auth := api.Http().Auth()
```

### Helpers

It returns an instance of the [HttpHelpers](./http-helpers.md).

```go
helpers := api.Http().Helpers()
```

### HttpRouter

It returns an instance of [HttpRouterApi](./http-router-api.md).

```go
httpRouter := api.Http().HttpRouter()
```

### Middlewares

Returns the built-in [middlewares](./http-router-api.md#middlewares).

```go
middlewares := api.Http().Middlewares()
```

Below are built-in middlewares available in the `Middlewares` instance:

```go
middlewares.AdminAuth() // It returns a middleware that checks if the user is authenticated.
middlewares.CacheResponse() // It returns a middleware that caches the response.
```


### HttpResponse

Returns an instance of [HttpResponse](./http-response.md).

```go
httpResponse := api.Http().HttpResponse()
```

### VueRouter

It returns an instance of [VueRouterApi](./vue-router-api.md).

```go
vueRouter := api.Http().VueRouter()
```

### VueResponse

Returns an instance of [VueResponse](./vue-response.md).

```go
vueResponse := api.Http().VueResponse()
```

### GetDevice

Get the device information from the http request. It returns and instance of [ClientDevice](./client-device.md) and an `error`.

```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    // other logic...
    device, err := api.Http().GetDevice(r)
    if err != nil {
        // handle error
    }
    fmt.Println(device) // ClientDevice
}
```

### MuxVars

Returns a `map[string]string` of variables from the request path. Below is an example to get the value if `id` in the route path `/sessions/:id`

```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    // other logic...
    vars := api.Http().MuxVars(r) // map[string]string
    id := vars["id"]
    fmt.Println(id) // "1"
}
```

### GetAdminNavs

Returns a slice of [AdminNavList](#adminnavlist)

```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    // other logic...
    navList := api.Http().GetAdminNavs(r)
    fmt.Println(navList) // []AdminNavList
}
```

### GetPortalItems

Returns a slice of [PortalItem](#portalitem)

```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    // other logic...
    portalItems := api.Http().GetPortalItems(r)
    fmt.Println(portalItems) // []PortalItem
}
```

## Admin Nav List {#adminnavlist}

`AdminNavList` is a list of [items](#adminnavitem) for the admin navigation. It has the following fields:

```go
type AdminNavList struct {
	Label string         `json:"label"`
	Items []AdminNavItem `json:"items"`
}
```

## Admin Nav Item {#adminnavitem}

`AdminNavItem` is an item for the admin navigation. It has the following fields:

```go
type AdminNavItem struct {
	Category       INavCategory      `json:"category"`
	Label          string            `json:"label"`
	VueRouteName   string            `json:"route_name"`
	VueRoutePath   string            `json:"route_path"`
	VueRouteParams map[string]string `json:"route_params"`
}
```

## Portal Item {#portalitem}

`PortalItem` is an item for the portal navigation. It has the following fields:

```go
type PortalItem struct {
	IconUri        string            `json:"icon_uri"`
	Label          string            `json:"label"`
	VueRouteName   string            `json:"route_name"`
	VueRoutePath   string            `json:"route_path"`
	VueRouteParams map[string]string `json:"route_params"`
}
```
