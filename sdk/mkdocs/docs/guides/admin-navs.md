# Admin Navs

Admin navs are links to the [admin routes](./basic-routing.md#admin-routes). Therefore, to display an item in the admin panel, one must already have defined an admin route.

To add an admin nav, we will use the [VueRouter.AdminNavsFunc](../api/vue-router.md#adminnavsfunc) method.

```go
api.Http().VueRouter().AdminNavsFunc(func(r *http.Request) []sdkhttp.VueAdminNav {
    adminNav := sdkhttp.VueAdminNav{
        Category:    sdkhttp.NavCategorySystem,
        Label:       "Welcome",
        RouteName:   "admin.welcome",
        RouteParams: map[string]string{"name": "Jhon"},
    }
    return []sdkhttp.VueAdminNav{adminNav}
})
```

An admin nav is defined by the following properties:

## Category
The admin navs are categorized to organize them according to their functionality. Although you can assign the navigation items to any category, it's better to put them to the closest relevant category for ease of navigation.

The `Category` field is defined by the `sdkhttp.NavCategory` type which is an enum with the following values:

```go
sdkhttp.NavCategorySystem   = "system"
sdkhttp.NavCategoryPayments = "payments"
sdkhttp.NavCategoryNetwork  = "network"
sdkhttp.NavCategoryThemes   = "themes"
sdkhttp.NavCategoryTools    = "tools"
```

## Label
This is the text label that will be displayed in the admin panel.

## RouteName
This is the name of the [admin route](./basic-routing.md#admin-routes) that the item will link to.

## RouteParams
This is a map of parameters that will be passed to the [admin route](./basic-routing.md#admin-routes). The parameters are used to build the URL of the route.
