# Saving and Retrieving Data

## 1. Saving Data

To save your plugin data like plugin settings, configuration and statistics, use the `Get` and `Save` methods from the [ConfigApi.Custom](../api/config-api.md#custom):

```go
type MyPluginConfig struct {
    MySetting       string  `json:"my_setting"`
    OtherSetting    int     `json:"other_setting"`
}

my_key := "my_config"
my_config := MyPluginConfig{
    MySetting:      "my_value",
    OtherSetting:   123,
}

err := api.Config().Custom(config_key).Save(my_config)
```

Plugin configuration is separated into different keys for ease of management. The data must be serializable to JSON.

## 2. Retreiving Data

To get your plugin data for a specific key, use the [ConfigApi.Custom.Get](../api/config-api.md#custom) method:
```go

my_config, err := api.Config().Custom(my_key).Get()
```


