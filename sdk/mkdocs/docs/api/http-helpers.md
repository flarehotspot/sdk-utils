# IHttpHelpers

`IHttpHelpers` is a set of helper methods that can be used in your http handlers to perform common tasks such as getting the client device, fetching client-side assets, translating messages, and more. It is accessible using the [IHttpApi.Helpers](./http-api.md#helpers) method:

```go
helprs := api.Http().Helpers()
```

## Methods

Below are the methods available in the `IHttpHelpers`:

### AssetPath

Returns the URI path of a manifest index filename.
For example to get the uri path of the file defined `resources/assets/css/style.css`:

```go
uri := api.Http().Helpers().AssetPath("css/style.css")
fmt.Println(uri) // /plugins/your-plugin-id/0.0.1/assets/css/style.css
```

### ResourcePath

Returns the URI path of a static file in `resources/assets` directory from your plugin.
For example to get the uri path of the file in `resources/assets/css/style.css`:

```go
uri := api.Http().Helpers().AssetPath("css/style.css")
fmt.Println(uri) // /plugins/your-plugin-id/0.0.1/assets/css/style.css
```

### AdsView

TODO: implement advertisements feature

### CsrfHtmlTag

Returns the CSRF HTML input tag as plain `string` to be used in HTML forms.

### Translate

Alias to [IPluginApi.Translate](./plugin-api.md#translate) method.

### UrlForRoute

Alias to [IHttpRouterApi.UrlForRoute](./http-router-api.md#urlforroute) method.

### UrlForPkgRoute

Alias to [IHttpRouterApi.UrlForPkgRoute](./http-router-api.md#urlforpkgroute) method.
