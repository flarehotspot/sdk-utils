# Portal Items

[Portal items](../api/http-api.md#portalitem) are the items displayed in the captive portal.

```
TODO: add image here
```

A portal item is just a link to an existing [portal route](./basic-routing.md#portal-routes). Therefore, to display an item in the captive portal, one must already have defined a portal route.

To register a portal item to be displayed in the captive portal, we will use the [VueRouter.PortalItemsFunc](../api/vue-router.md#portalitemsfunc) method:

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
