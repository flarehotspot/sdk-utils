# ClientDevice

## 1. Events {#events}

Events are emitted to the client device in the browser.

You can emit an event to a user account using the [ClientDevice.Emit](#emit) method like so:

```go
func (w http.ResponseWriter, r *http.Request) {
    clnt, _ := api.Http().Helpers().GetClientDevice(r)
    evt := "some_event"
    data := map[string]interface{}{"key": "value"}
    clnt.Emit(evt, data)
}
```

You can listen to this events in the browser using the [$flare.events](./flare-variable.md#flare-events) like so:

```js
$flare.events.on("some_event", function(res) {
    console.log("An event occured: ", res.data);
});
```

The avaiable events geenerated by the system are:

| Event | Description
|--|--
| `session:connected` | Emitted when a client device is connected to the internet.
| `session:disconnected` | Emitted when a client device is disconnected from the internet.
