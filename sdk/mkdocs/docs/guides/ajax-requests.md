# Ajax Requests
Ajax is a technique that allows a web page to be updated asynchronously by exchanging small amounts of data with the server behind the scenes. This means that it is possible to update parts of a web page, without reloading the entire page.

## 1. GET Request {#get-request}
To make a `GET` request, we need an existing [GET HTTP Route](../api/http-router-api.md#get) and [handler function](#handler-function).

```go
router := api.Http().HttpRouter().PluginRouter()
router.Get("/welcome/:name", func(w http.ResponseWriter, r *http.Request) {
    // Handle the request
    vars := api.Http().MuxVars(r) // map[string]string
    name := vars["name"]
    data := map[string]string{
        "name": name
    }

    // send JSON http response
    api.Http().HttpResponse().Json(w, data, http.StatusOK)

}).Name("payments.options") // set the route name
```

## 2. POST Request {#post-request}

## 3. Handler Function {#handler-function}


