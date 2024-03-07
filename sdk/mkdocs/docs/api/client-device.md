# ClientDevice

## 4. Events {#events}
Events are emitted to the user accounts via SSE (Server-Sent Events) in the browser.
You can listen to this events in the browser using the [$flare.events](./flare-variable.md#flare-events) object like so:
```js
$flare.events.addEventListener("some_event", function(res) {
    console.log("An event occured: ", res.data);
});
```
