# The `$flare` global variable
The `$flare` variable is a global variable in the browser that contains helper functions to work with the Flare API.

## $flare.http {#flare-http}
The `$flare.http` object is used to make AJAX requests. It contains two methods, namely `get` and `post`.

### $flare.http.get {#flare-http-get}
The `$flare.http.get` method accepts two arguments, the first argument is the URL to send the form data to, and the second argument is the query params.
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
The `$flare.http.post` method accepts two arguments, the first argument is the URL to send the form data to, and the second argument is the form data.

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

## $flare.events {#flare-events}
The `$flare.events` object is used to listen to events emitted by the server via [Server-Sent Events](https://www.w3schools.com/html/html5_serversentevents.asp). Below is an example of how to listen to an event:
```js
$flare.events.addEventListener("session:connected", function(res) {
    console.log("Session connected: ", res.data);
});
```

!!!note
    Flare Hotspot SDK already includes a polyfill for the `EventSource` object for browsers that don't support Server-Sent Events.
