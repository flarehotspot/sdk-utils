# HttpResponse
The `HttpResponse` has utility functions which can be used to send html, json, and file response to the client.

## 1. HttpResponse Methods {#httpresponse-methods}

### PortalView
This method is used to render views as plain html from `resources/views/portal` directory in your plugin.
For example if you have a view in `resources/views/portal/index.html`,
then you can render it with:

```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{
        "title": "Dashboard",
    }
    api.Http().HttpResponse().PortalView(w, r, "dashboard/index.html", data)
}
```

It uses the file `resources/views/portal/layout.html` as the [layout](#layout-view). You must create this file in order to use the `PortalView` method.
The view has access to the [HttpHelpers](./http-helpers.md) instance. Below is an example of how to access the data and helpers in the view:

```html title="resources/views/portal/index.html"
<a href='<% .Helpers.UrlForRoute "dashboard" %>'>
    <% .Data.title %>
</a>
```

### AdminView
This method is very similar to [PortalView](#portalview) but it is used to render views from `resources/views/admin` directory in your plugin. It also uses the file `resources/views/admin/layout.html` as the [layout](#layout-view). You must create this file in order to use the `AdminView` method.

Below is an example of how to render a view from the admin directory:

```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{
        "title": "Dashboard",
    }
    api.Http().HttpResponse().PortalView(w, r, "dashboard/index.html", data)
}
```

```html title="resources/views/admin/index.html"
<a href='<% .Helpers.UrlForRoute "dashboard" %>'>
    <% .Data.title %>
</a>
```

### View
This method is similar to [PortalView](#portalview) and [AdminView](#adminview) but it is used to render any views from the `resources/views` directory in your plugin as plain html. It does not use any layout view. It also has access to the [HttpHelpers](./http-helpers.md) instance and data.

### File
This method is used to render text and asset files from the `resources` directory. Just like [PortalView](#portalview) and [AdminView](#adminview) methods, the template files have access to the [HttpHelpers](./http-helpers.md) instance and data. The response header's `Content-Type` will be automatically derived from the filename. Below is an example of how to render a file:

```js title="resources/assets/js/app.js"
var url = '<% .Helpers.UrlForRoute .Data.title %>';
console.log(url);
```

```go title="main.go"
data := map[string]string{}
    "title": "Dashboard",
}
api.Http().HttpResponse().File(w, r, "assets/js/app.js", data)
```

### Json
This method is used to send json response to the client. Below is an example of how to send json response:
```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
        "title": "Dashboard",
    }
    api.Http().HttpResponse().Json(w, data, http.StatusOK)
}
```

## 2. Template Parsing {#template-parsing}
The views are parsed using the [html/template](https://pkg.go.dev/html/template) package. But instead of using `{{ }}` as delimiters, we are using `<% %>` as delimiters. This is to avoid conflicts with the `{{ }}` delimiters used in the frontend framework.

## 3. Layout View {#layout-view}
In order to use [HttpResponse.PortalView](#portalview) and [HttpResponse.AdminView](#adminview) methods, a layout view must be created first.
The portal layout view must be created in `resources/views/portal/layout.html` and the admin layout view must be created in `resources/views/admin/layout.html` inside your plugin directory.

The layout view is used to define the common structure of the view. For example, the layout view can be used to define the header, footer, and sidebar of the page. Below is an example of how to define the layout view:

```html title="resources/views/portal/layout.html"
<!doctype html>
<html lang="en">
    <head></head>
    <body>
        <h1>Portal Layout</h1>
        <div class="container"> <% .ContentHtml %> </div>
    </body>
</html>
```
