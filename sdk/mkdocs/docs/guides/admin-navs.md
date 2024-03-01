# Admin Navs

[Admin navs](../api/http-api.md#adminnavitem) are navigation links that points to an [admin route](./routes-and-links.md#admin-routes). Therefore, to display an item in the admin panel, one must already have defined an admin route.

To add an admin nav, we will use the [VueRouterApi.AdminNavsFunc](../api/vue-router-api.md#adminnavsfunc) method.

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

## Admin Nav Definition
An admin nav is defined by the following fields:

### Category
The admin navs are categorized to organize them according to their functionality. Although you can assign the navigation items to any category, it's better to put them to the closest relevant category for ease of navigation.

The `Category` field is defined by the `sdkhttp.NavCategory` type which is an enum with the following values:

```go
sdkhttp.NavCategorySystem   = "system"
sdkhttp.NavCategoryPayments = "payments"
sdkhttp.NavCategoryNetwork  = "network"
sdkhttp.NavCategoryThemes   = "themes"
sdkhttp.NavCategoryTools    = "tools"
```

### Label
This is the text label that will be displayed in the admin panel.

### RouteName
This is the name of the [admin route](./routes-and-links.md#admin-routes) that the item will link to.

### RouteParams
This is a map of parameters that will be passed to the [admin route](./routes-and-links.md#admin-routes). The parameters are used to build the URL of the route.
