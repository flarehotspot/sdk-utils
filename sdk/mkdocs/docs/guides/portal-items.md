# Portal Items

[Portal items](../api/http-api.md#portalitem) are the items displayed in the captive portal.

```
TODO: add image here
```

A portal item is just a link to an existing [portal route](./routes-and-links.md#portal-routes). Therefore, to display an item in the captive portal, one must already have defined a portal route.

To register a portal item to be displayed in the captive portal, we will use the [VueRouterApi.PortalItemsFunc](../api/vue-router-api.md#portalitemsfunc) method:

```go
api.Http().VueRouter().PortalItemsFunc(func(r *http.Request) []sdkhttp.VuePortalItem {
    portalItem := sdkhttp.VuePortalItem{
        Label:       "Welcome",
        IconPath:    "icons/welcome.png",
        RouteName:   "portal.welcome",
        RouteParams: map[string]string{"name": "Jhon"},
    }
    return []sdkhttp.VuePortalItem{portalItem}
})
```

## Portal Item Definition
A portal item is defined by the following fields:

### Label
This is the text label that will be displayed in the captive portal.

### IconPath
This is the path to the icon image that will be displayed next to the label. Icon paths are searched in `resources/assets` directory.

### RouteName
This is the [route name](./routes-and-links.md#routename) field of the portal route that the item will link to.

### RouteParams
This is a map of [parameters](./routes-and-links.md#route-params) that will be passed to the [portal route](./routes-and-links.md#portal-routes) to form the link.
