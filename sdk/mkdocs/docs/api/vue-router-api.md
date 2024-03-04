# VueRouterApi

This is the API for the frontend router.

## VueRouterApi Methods

### RegisterPortalRoutes

This method is used to register the [portal routes](#portalroute).

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

This method is used to register the [admin routes](#adminroute).
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

### AdminNavsFunc

### VueRouteName

### VueRoutePath

### VuePkgRoutePath

## PortalRoute

## AdminRoute
