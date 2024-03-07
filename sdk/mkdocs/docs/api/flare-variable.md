# The `$flare` global variable
The `$flare` variable is a global variable in the browser that contains helper functions to work with the Flare Hotspot API.

### $flare.http.get {#flare-http-get}
The `$flare.http.get` method is used to perform a `GET` AJAX request. It accepts two arguments, the first argument is the URL to send the form data to, and the second argument is the query params.
```js
var queryParams = {amount: 100};

$flare.http.get('/path/to/handler', queryParams)
    .then(function(response){
        console.log(response);
    })
    .catch(function(error){
        console.log(error);
    });
```

### $flare.http.post {#flare-http-post}
The `$flare.http.post` method is used to perform a `POST` AJAX request. It accepts two arguments, the first argument is the URL to send the form data to, and the second argument is the form data.

```js
var formData = {amount: 100};

$flare.http.post('/path/to/handler', formData)
    .then(function(response){
        console.log(response);
    })
    .catch(function(error){
        console.log(error);
    });
```

!!!warning "Important"
    You must use [VueResponse](./vue-response.md) in the server side to perform http resposes for both the [$flare.http.get](#flare-http-get) and [$flare.http.post](#flare-http-post) methods.

## $flare.vueLazyLoad
The `$flare.vueLazyLoad` method is used to lazy load vue components.

```js
var component = '<% .Helpers.VueComponentPath "sample-child.vue" %>';
var lazyComponent = $flare.vueLazyLoad(component);

var app = new Vue({
    el: '#app',
    components: {
        'sample-child': lazyComponent
    }
});
```

## $flare.events {#flare-events}
The `$flare.events` object is used to listen to events emitted by the server via [Server-Sent Events](https://www.w3schools.com/html/html5_serversentevents.asp). Below is an example of how to listen to an event:
```js
$flare.events.addEventListener("session:connected", function(res) {
    console.log("Session connected: ", res.data);
});
```

See the user account events in the [AccountsApi](./accounts-api.md#events) documentation.

See the client device events in the [ClientDevice](./client-device.md#events) documentation.

!!!note
    Flare Hotspot SDK already includes a polyfill for the `EventSource` object for browsers that don't support Server-Sent Events.

