# VueResponse

## VueResponse Methods

The methods within `VueResponse` are used to send http response to the browser that's using the [$flare.http.get](./flare-variable.md#flare-http-get) or [$flare.http.post](./flare-variable.md#flare-http-post) helper methods.

### SetFlashMsg

This method is to send a toast message to the client. This does not send a response to the client.
Thus, it must be called along with [Json](#json), [Redirect](#redirect) or
[RedirectToPortal](#redirecttoportal) methods to send the response to the client.
See the [Json](#json) documentation for sending json data.

```go
data := nil
res := api.Http().VueResponse()
res.SetFlashMsg("success", "Hello, World!")
res.Json(w, data, http.StatusOK)
```

### SendFlashMsg

This method is similar to [SetFlashMsg](#setflashmsg) but it sends the message to the client immediately. Its only use is to send a toast message to the client.

```go
res := api.Http().VueResponse()
res.SendFlashMsg(w, "success", "Hello, World!", http.StatusOK)
```

### Json

Used to send a JSON response to the client. The data can be any type as long as it can be marshalled to JSON.

```go
res := api.Http().VueResponse()
res.Json(w, data, http.StatusOK)
```

### Redirect

This methods redirects the client to another registered [Vue Route](../guides/routes-and-links.md) in your plugin.

```go
res := api.Http().VueResponse()
res.Redirect(w, "route-name")
```

### RedirectToPortal

This method redirects the client to the index page of the captive portal.

```go
res := api.Http().VueResponse()
res.RedirectToPortal(w)
```

### Error

This metohd sends an error toast message to the client.

```go
res := api.Http().VueResponse()
res.Error(w, "An error occurred", http.StatusInternalServerError)
```
