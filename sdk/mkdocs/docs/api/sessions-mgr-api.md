# SessionsMgrApi

The `SessionsMgrApi` contains methods to manage the [ClientDevice](./client-device.md) sessions.

## SessionsMgrApi Methods

### Connect

This method will connect the client device to the internet if the client device has available session to consume. Otherwise, it will return an error.

```go
func (w http.ResponseWriter, r *http.Request) {
    clnt, _ := api.Http().GetClientDevice(r)
    err = api.SessionsMgr().Connect(clnt)
}
```

### Disconnect

This method will disconnect the client device from the internet. It will also pause the current running session of the client device.

```go
func (w http.ResponseWriter, r *http.Request) {
    clnt, _ := api.Http().GetClientDevice(r)
    err = api.SessionsMgr().Disconnect(clnt)
}
```
