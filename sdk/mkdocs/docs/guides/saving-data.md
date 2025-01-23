# Saving Data

## Saving Data

To save the plugin data like plugin settings, configuration and statistics, we use the  [IPluginCfgApi.Write](../api/config-api.md#write):

```go
import "encoding/json"
// ...

type MyPluginConfig struct {
    MySetting       string  `json:"my_setting"`
    OtherSetting    int     `json:"other_setting"`
}

myConfig := MyPluginConfig{
    MySetting:      "my_value",
    OtherSetting:   123,
}

data, err := json.Marshal(myConfig)
if err != nil {
    // handle error
}

err := api.Config().Plugin().Write("my_key", data)
```

Plugin configuration is separated into different keys for ease of management. The data must be serializable to JSON.

## Retreiving Data

To get the plugin data for a specific key, use the [IPluginCfgApi.Read](../api/config-api.md#read) method:
```go
import "encoding/json"
// ...

data, err := api.Config().Plugin().Read("my_key")
if err != nil {
    // handle error
}

var myConfig MyPluginConfig
if err := json.Unmarshal(data, &myConfig); err != nil {
    // handle error
}

fmt.Println(myConfig) // {MySetting: "my_value", OtherSetting: 123}
```


