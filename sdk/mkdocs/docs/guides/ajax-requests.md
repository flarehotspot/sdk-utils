# AJAX Requests
**AJAX** is a technique that allows a web page to be updated asynchronously by exchanging small amounts of data with the server behind the scenes. This means that it is possible to update parts of a web page, without reloading the entire page.

!!! Note
    We only support two (2) HTTP methods - `GET` and `POST` request. This is to simplify the front-end code and to be able to support wider range of browsers.

## 1. GET Request {#get-request}
To make a `GET` request, we need an existing [GET route](../api/http-router-api.md#get) and [handler function](#handler-function).

Below is an example of a `GET` request to the `/welcome` route.
This route and handler simply sends a JSON response with a message `Hello World!`.

```go title="routes.go"
router := api.Http().HttpRouter().PluginRouter()
router.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
        "message": "Hello World!"
    }
    // send JSON http response
    api.Http().HttpResponse().Json(w, data, http.StatusOK)

}).Name("welcome") // set the route name
```

Now, we need to create our [vue component](./vue-components.md) to display the message from our handler:

```html title="Welcome.vue"
<template>
    <p>{{ message }}</p>
</template>

<script>
define(function(){
    return {
        template: template,
        data: function() {
            return {
                message: ""
            }
        },
        mounted: function() {
            var self = this;
            $flare.http.get('<% .Helpers.UrlForRoute "welcome" %>')
            .then(function(data) {
                self.message = data.message;
            });
        }
    }
})
</script>
```

In this example, our vue component simply fetches the message from the http route named `welcome` and displays the message in the template.

## 2. POST Request {#post-request}

Performing a `POST` request is very similar to `GET` request, except we use the [$flare.http.post](../api/flare-variable.md#flare-http-post) method.

Take a look at [Form Submission](./form-submission.md) for an example on how to make a `POST` request.

## 3. Handler Function {#handler-function}

A handler function is simply a function that handles the http request. It accepts two arguments, the `http.ResponseWriter` and `*http.Request`. For example:

```go title="main.go"
func (w http.ResponseWriter, r *http.Request) {
    // handle the request
}
```

To respond to AJAX requests, we use the methods available in [VueResponse](../api/vue-response.md). For example, to send a JSON response:

```go title="main.go"
func (w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
        "message": "Hello World!"
    }
    // send JSON http response
    api.Http().HttpResponse().Json(w, data, http.StatusOK)
}
```
