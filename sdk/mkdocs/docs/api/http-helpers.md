# HttpHelpers

## GetClientDevice
Get the [client device](./client-device.md) info from the http request:
```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    device := api.Http().Helpers().GetClientDevice(r)
    fmt.Println(device) // ClientDevice
}
```

## Translate
Alias to [PluginApi.Translate](./plugin-api.md#translate) method.

## AssetPath
Returns the URI path of a static file in `resources/assets` directory from your plugin.
For example to get the uri path of the file in `resources/assets/css/style.css`:
```go
uri := api.Http().Helpers().AssetPath("css/style.css")
fmt.Println(uri) // /plugins/your-plugin-id/0.0.1/assets/css/style.css
```

## AssetWithHelpersPath
Similar to [AssetPath](#assetpath), but the assets are pre-processed with access to the `HttpHelper` instance.
For example, if you want to have an asset that points to a certain route in your plugin, you can do:
```go title="main.go"
uri := api.Http().Helpers().AssetWithHelpersPath("js/script.js")
fmt.Println(uri) // /plugins/your-plugin-id/0.0.1/assets/js/script.js
```

```js title="resources/assets/js/script.js"
var url = '<% .Helpers.VueRoutePath "some.routename" %>';
console.log(url);
```

## VueComponentPath
Returns the URI path of a file in `resources/components` directory from your plugin.
The vue component is parsed using text/template go module (using `<%` and `%>` delimiters) and has access to `<% .Helpers %>` object. This is often used to [lazy load](../guides/vue-components.md#lazy-loading-components) vue components.

```go
uri := api.Http().Helpers().VueComponentPath("sample-component.vue")
fmt.Println(uri) // /plugins/your-plugin-id/0.0.1/components/sample-component.vue
```

## VueRoutePath
Alias to [VueRouterApi.VueRouterPath](./vue-router-api.md#vueroutepath) method.
