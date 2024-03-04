# HttpResponse

The `HttpResponse` has utility functions which can be used to send html, json, and file response to the client.

## HttpResponse Methods

### PortalView

This method is used to render views from `/resources/views/portal` directory in your plugin.
For example if you have a view in `/resources/views/portal/index.html`,
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

It uses the file `/resources/views/portal/layout.html` as the [layout](#layout-view).
The view has access to the [HttpHelpers](./http-helpers.md) instance. Below is an example of how to access the data and helpers in the view:

```html title="resources/views/portal/index.html"
<a href='<% .Helpers.UrlForRoute "dashboard" %>'>
    <% .Data.title %>
</a>
```

### AdminView

This method is very similar to [PortalView](#portalview) but it is used to render views from `/resources/views/admin` directory in your plugin. It also uses the file `/resources/views/admin/layout.html` as the [layout](#layout-view).
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

## Layout View

The layout view is a wrapper around the view which is used to render the view. It is used to define the structure of the view. The layout view is used to define the common structure of the view. For example, the layout view can be used to define the header, footer, and sidebar of the view. Below is an example of how to define the layout view:
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
