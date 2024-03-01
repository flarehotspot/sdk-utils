# HttpHelpers

`HttpHelpers` is a set of helper methods that can be used in your http handlers to perform common tasks such as getting the client device, fetching client-side assets, translating messages, and more. It is accessible using the [Http.Helpers](./http-api.md#helpers) method:

```go
helprs := api.Http().Helpers()
```

It can also be accessed in the views, assets and vue components using `<% .Helpers %>`.

## Methods

Below are the methods available in the `HttpHelpers`:

### AssetPath
Returns the URI path of a static file in `resources/assets` directory from your plugin.
For example to get the uri path of the file in `resources/assets/css/style.css`:
```go
uri := api.Http().Helpers().AssetPath("css/style.css")
fmt.Println(uri) // /plugins/your-plugin-id/0.0.1/assets/css/style.css
```

### AssetWithHelpersPath
Similar to [AssetPath](#assetpath), but the assets are pre-processed with access to the `HttpHelpers` instance.
For example, if you want to have an asset that points to a certain route in your plugin, you can do:
```go title="main.go"
uri := api.Http().Helpers().AssetWithHelpersPath("js/script.js")
fmt.Println(uri) // /plugins/your-plugin-id/0.0.1/assets/js/script.js
```

```js title="resources/assets/js/script.js"
var url = '<% .Helpers.VueRoutePath "some.routename" %>';
console.log(url);
```

### VueComponentPath
Returns the URI path of a file in `resources/components` directory from your plugin.
The vue component is parsed using text/template go module (using `<%` and `%>` delimiters) and has access to `<% .Helpers %>` object. This is often used to [lazy load](../guides/vue-components.md#lazy-loading-components) vue components.

```go
uri := api.Http().Helpers().VueComponentPath("sample-component.vue")
fmt.Println(uri) // /plugins/your-plugin-id/0.0.1/components/sample-component.vue
```

### VueRoutePath
Alias to [VueRouterApi.VueRouterPath](./vue-router-api.md#vueroutepath) method.

### AddView
TODO: implement advertisements feature

### Translate
Alias to [PluginApi.Translate](./plugin-api.md#translate) method.

### MuxRouteName
Alias to [HttpRouterApi.MuxRouteName](./http-router-api.md#muxroutename) method.

### UrlForMuxRoute
Alias to [HttpRouterApi.UrlForMuxRoute](./http-router-api.md#urlformuxroute) method.

### UrlForRoute
Alias to [HttpRouterApi.UrlForRoute](./http-router-api#urlforroute) method.

### UrlForPkgRoute
Alias to [HttpRouterApi.UrlForPkgRoute](./http-router-api#urlforpkgroute) method.

### VueRouteName
Alias to [VueRouterApi.VueRouteName](./vue-router-api.md#vueroutename) method.

### VueRoutePath
Alias to [VueRouterApi.VueRoutePath](./vue-router-api.md#vueroutepath) method.

### VuePkgRoutePath
Alias to [VueRouterApi.VuePkgRoutePath](./vue-router-api.md#vuepkgroutepath) method.
