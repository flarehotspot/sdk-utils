# SessionsMgrApi

The `SessionsMgrApi` contains methods to manage the [ClientDevice](./client-device.md) sessions.

## SessionsMgrApi Methods

### Connect

This method will connect the client device to the internet if the client device has available [ClientSession](./client-session.md) to consume.
It takes a [context](https://gobyexample.com/context), a [ClientDevice](./client-device.md), and a notification `string` as parameters.

```go
func (w http.ResponseWriter, r *http.Request) {
    clnt, _ := api.Http().GetClientDevice(r)
    err = api.SessionsMgr().Connect(r.Context(), clnt, "You are now connected to internet.")
}
```

### Disconnect

This method will disconnect the client device from the internet. It will also pause the current running [ClientSession](./client-session.md) of the client device. It takes a [context](https://gobyexample.com/context), a [ClientDevice](./client-device.md) and a notification `string` as parameters.

```go
func (w http.ResponseWriter, r *http.Request) {
    clnt, _ := api.Http().GetClientDevice(r)
    err = api.SessionsMgr().Disconnect(r.Context(), clnt, "You are now disconnected to internet.")
}
```

### IsConnected

Returns `true` if the [ClientDevice](./client-device.md) is connected to the internet, otherwise `false`.

```go
func (w http.ResponseWriter, r *http.Request) {
    clnt, _ := api.Http().GetClientDevice(r)
    isConnected, err = api.SessionsMgr().IsConnected(clnt)
}
```

### CreateSession

It creates a [ClientSession](./client-session.md) for the [ClientDevice](./client-device.md). It takes the following arguments:

- `context.Context`
- `int64` - the [ClientDevice](./client-device.md) ID
- `uint8` - the [type of session](./client-session.md#type) to create
- `uint` - the duration of the session in seconds, applicable only for `time` and `time_or_data` session types
- `float64` - the data in mega bytes, applicable only for `data` and `time_or_data` session types
- `*uint` - the expiration in days after the session is started, on top of the duration in seconds
- `int` - the download speed of the session in megabits per second (mbps)
- `int` - the upload speed of the session in megabits per second (mbps)
- `bool` - whether to use the global download and upload speed limit. If `true`, it ignores the download and upload speed arguments

Below is an example of how to use the `CreateSession` method:

```go
func (w http.ResponseWriter, r *http.Request) {
    secs := 60          // 1 minute
    mb := 100.0         // 100 MB
    sessionType := 0    // 0 = time, 1 = data, 2 = time_or_data
    expireDays := 30    // 30 days
    downMbits := 5      // 5 mbps
    uploadMbits := 3    // 3 mbps

    clnt, _ := api.Http().GetClientDevice(r)
    err := api.SessionsMgr().CreateSession(
        r.Context(),
        clnt.Id(),
        sessionType,
        secs,
        mb,
        &expireDays,
        downMbits,
        upMbits,
        false,
    )
}
```

### CurrSession

This is the current running [ClientSession](./client-session.md) of the [ClientDevice](./client-device.md).

```go
func (w http.ResponseWriter, r *http.Request) {
    clnt, _ := api.Http().GetClientDevice(r)
    session, ok = api.SessionsMgr().CurrSession(clnt)
}
```

### GetSession

Returns any available [ClientSession](./client-session.md) for the given [ClientDevice](./client-device.md) ID. This may include the current running session or any paused session.

```go
func (w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    clnt, _ := api.Http().GetClientDevice(r)
    session, err = api.SessionsMgr().GetSession(ctx, clnt)
}
```

### RegisterSessionProvider

Used to register a new session provider function. The function accepts a `context.Contxt` and a [ClientDevice](./client-device.md) parameters and it session should return an instance of [SessionSource](./session-source.md) and an `error` if any.

The example below provides a `time` [session type](./client-session.md#session-types) for every client device:

```go
type RemoteSession struct {
    mu          sync.RWMutex
    timeSeconds uint
}

func (rs *RemoteSession) Data() sdkconnmgr.SessionData {
    return sdkconnmgr.SessionData{
        Provider: "remote-session",
        Type: 0,
        TimeSecs: rs.timeSeconds,
    }
}

func (rs *RemoteSession) Save(ctx context.Context, data sdkconnmgr.SessionData) error {
    self.mu.Lock()
    defer self.mu.Unlock()
    // implemnt save logic here
    rs.timeSeconds = data.TimeSecs
    return nil
}

func (rs *RemoteSession) Reload(ctx context.Context) (sdkconnmgr.SessionData, error {
    // implement reload logic here
    return sdkconnmgr.SessionData{
        Provider: "remote-session",
        Type: 0,
        TimeSecs: rs.timeSeconds,
    }, nil
}

api.SessionsMgr().RegisterSessionProvider(func(ctx context.Context, clnt *sdkconnmgr.ClientDevice) (sdkconnmgr.SessionSource, error) {
    // give every client device a 1 minute session
    return &RemoteSession{
        timeSeconds: 60,
    }, nil
})
```
