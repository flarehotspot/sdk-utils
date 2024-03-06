# VueRouterApi

This is the API for the frontend router.

## 1. VueRouterApi Methods {#vuerouterapi-methods}

### RegisterPortalRoutes

This method is used to register the [portal routes](#portalroute). It is a *variadic function* that accepts multiple instances of `sdkhttp.VuePortalRoute`.

```go
api.Http().VueRouter().RegisterPortalRoutes(sdkhttp.VuePortalRoute{
    RouteName:   "insert-coin",
    RoutePath:   "/coinslot/:id/insert-coin",
    Component:   "InsertCoin.vue",
    HandlerFunc: func (w http.ResponseWriter, r *http.Request) {
        // Do something
    },
})
```

### RegisterAdminRoutes

This method is used to register the [admin routes](#adminroute). It is a *variadic function* that accepts multiple instances of `sdkhttp.VueAdminRoute`.

```go
api.Http().VueRouter().RegisterAdminRoutes(sdkhttp.VueAdminRoute{
    RouteName:   "admin-dashboard",
    RoutePath:   "/admin/dashboard",
    Component:   "AdminDashboard.vue",
    HandlerFunc: func (w http.ResponseWriter, r *http.Request) {
        // Do something
    },
})
```

### PortalItemsFunc

This method is used to show items in the captive portal.
The items are a slice of `sdkhttp.VuePortalItem`.
This function is called every time the captive portal is loaded.
See [Portal Items](../guides/portal-items.md) for example.

### AdminNavsFunc

This method is used to add items to the admin navigation.
The items are a slice of `sdkhttp.VueAdminNav`.
See [Admin Navs](../guides/admin-navs.md) for examples.

### VueRouteName

This method returns the vue route name that can be used for `router-link` vue component:

```go
data := map[string]string{
    "VueRouteName": api.Http().VueRouter().VueRouteName("insert-coin"),
}
```

```html
<router-link :to="{name: <% .Data.VueRouteName %>}">
```

### VueRoutePath

This method returns the vue route path that can be used for `router-link` vue component:

```go
data := map[string]string{
    "VueRoutePath": api.Http().VueRouter().VueRoutePath("insert-coin"),
}
```

```html
<router-link :to="vueRoutePath">
```

### VuePkgRoutePath

This method returns the vue route path from third-party plugins that you can use on your own plugin.

```go
data := map[string]string{
    "VueRoutePath": api.Http().VueRouter().VuePkgRoutePath("com.third-party.plugin", "third-party-route"),
}
```

```html
<router-link :to="vueRoutePath">
```

## 2. PortalRoute {#portalroute}

A portal route is an instance of `sdkhttp.VuePortalRoute` that is used to register portal routes.

See [Portal Routes](../guides/routes-and-links.md#portal-routes).

## 3. AdminRoute {#adminroute}

An admin route is an instance of `sdkhttp.VueAdminRoute` that is used to register admin routes.

See [Admin Routes](../guides/routes-and-links.md#admin-routes).
