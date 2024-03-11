# ConfigApi

## 1. Application {#application}

The application configuration has the following fields:

Currency
: The currency used throughout the application. The default is `USD`.

Lang
: The language used throughout the application. The default is `en`. The supported languages are:

- `en` - English
- `es` - Spanish
- `id` - Indonesian
- `ms` - Malay

Secret
: The secret key used to sign the JWT tokens and other encryptions.

To get the application configuration, use the `Get` method.

```go
cfg, err := api.Config().Application().Get()
```

To modify the application configuration, use the `Save` method.

```go
err := api.Config().Application().Save(sdkcfg.AppCfg{
    Currency: "USD",
    Lang: "en",
    Secret: "xxxxxxxxxx"
})
```

## 2. Bandwidth {#bandwidth}

Bandwidth configuration is set per interface and has the following fields:

UseGlobal
: Whether to use the global bandwidth configuration. The default is `false`.

GlobalDownMbits
: The global download bandwidth in megabits per second. The default is `2`.

GlobalUpMbits
: The global upload bandwidth in megabits per second. The default is `2`.

UserDownMbits
: The download speed per user in megabits per second. The default is `2`.

UserUpMbits
: The upload speed per user in megabits per second. The default is `2`.

To get the bandwidth configuration of a network interface, use the `Get` method.

```go
cfg, err := api.Config().Bandwidth("eth0").Get()
```

To set the bandwidth configuration of a network interface, use the `Save` method.

```go
err := api.Config().Bandwidth("eth0").Save(sdkcfg.BandwdCfg{
    UseGlobal: true,
    GlobalDownMbits: 2,
    GlobalUpMbits: 2,
    UserDownMbits: 2,
    UserUpMbits: 2,
})
```

## 3. Plugin {#plugin}

The plugin configuration API is used to store custom configuration specific to the plugin you are developing. Using this API ensures that your custom plugin configuration can be migrated properly to a new system in case you want to flash new firmware or migrate to a new hardware.

To save your plugin configuration, use the `Save` method.

```go

type MyPluginCfg struct {
    Field1 string `json:"field1"`
    Field2 string `json:"field2"`
}

err := api.Config().Plugin("myplugin").Save(MyPluginCfg{
    Field1: "value1",
    Field2: "value2",
})
```

To get your plugin configuration, use the `Get` method.

```go
var myCfg MyPluginCfg

err := api.Config().Plugin("myplugin").Get(&myCfg)

fmt.Printf("Field1: %s, Field2: %s", myCfg.Field1, myCfg.Field2)
```
