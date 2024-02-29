# PluginApi

The `PluginApi` is the root interface of Flare Hotspot SDK. It provides access to methods used to manipulate system accounts, network devices, theme configuration, user sessions, payment system and more. Each plugin is provided with an instance of `PluginApi`.

When the plugin is first loaded into the system, the system looks for the `Init` function of the plugin's `main` package. The `PluginApi` object is then passed to the plugin's `Init` function. From here, you can start configuring the routes and components of your plugin. An example of a plugin's init function:

```go title="plugins/com.mydomain.myplugin/main.go"
package main

import (
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func main() {}

func Init(api sdkplugin.PluginApi) {
    // You can start using the SDK here.
    // You can configure your routes, define your plugin components
    // and register items in the portal and admin nav menu, and more.
}
```

## PluginApi Methods

The following are the available methods in `PluginApi`.

### Name
It returns the `name` field defined in `plugin.json`.

```go
name := api.Name()
fmt.Println(name) // "My Plugin"
```

### Pkg
It returns the `package` field defined in [plugin.json](../plugin-json/).
```go
pkg := api.Pkg()
fmt.Println(pkg) // "com.mydomain.myplugin"
```

### Version
It returns the `version` field defined in [plugin.json](../plugin-json/).
```go
version := api.Version()
fmt.Println(version) // "1.0.0"
```

### Description
It returns the `description` field defined in [plugin.json](../plugin-json/).
```go
description := api.Description()
fmt.Println(description) // "My plugin description"
```

### Dir
It returns the absolute path of the plugin's installtion directory.
```go
dir := api.Dir()
fmt.Println(dir) // "/path/to/com.mydomain.myplugin"
```

### Translate
It is a utility function used to convert a message key into a translated string. Example usage:
```go
msg := api.Translate("info", "payment_received", "amount", 1.00)
fmt.Println(msg) // "Payment received USD 1.0.0"
```

In this example, given that the [application](../api/config-api.md#application) language is set to `en`, the system will look for the file `resources/translations/en/info/payment_received.txt` inside your plugin directory. If the file is found, the system will use the contents of the file as the translation template.

Sometimes we want to put variables inside the translation message. In this example, we want to pass the `amount` as a paramenter to the message. We can do that by passing the amount param as key-value pairs to the `Translate` method. Internally, the param pairs are converted into a type `map[any]any`. To use the `amount` param in the translation file, we'll enclose it with `<%` and `%>` delimiters (with dot prefix). Therefore the content of `payment_received.txt` should be:
```go
Payment received: USD <% .amount %>
```

### Resource
It returns the absolute path of the file under the plugin's resource directory.
```go
resource := api.Resource("/my-resource.txt")
fmt.Println(resource) // "/path/to/com.mydomain.myplugin/resources/my-resource.txt"
```

### SqlDb
It returns [*sql.DB](http://go-database-sql.org/overview.html) instance which is used to query, insert, update and delete database entities.
```go
db := api.SqlDb()
fmt.Println(db) // *sql.DB
```

### Acct
It returns the [AccountsApi](./accounts-api.md) object which is used to access and modify the system admin accounts.
```go
acct := api.Acct()
fmt.Println(acct) // AccountsApi
```

### Http
It returns the [`HttpApi`](./http-api.md) object which is used to configure routes and serve HTTP requests.
```go
http := api.Http()
fmt.Println(http) // HttpApi
```

### Config
It returns the [`ConfigApi`](./config-api.md) object which is used to access and modify the system configuration.
```go
config := api.Config()
fmt.Println(config) // ConfigApi
```

### Payments
It return the [`PaymentsApi`](../payments-api/) object which is used to create payment options or create system transactions.
```go
payments := api.Payments()
fmt.Println(payments) // PaymentsApi
```

### InAppPurchases
It returns the [`InAppPurchasesApi`](../in-app-purchases-api/) object which is used to create and manage in-app purchases.
```go
inAppPurchases := api.InAppPurchases()
fmt.Println(inAppPurchases) // InAppPurchasesApi
```

### Ads
It returns the [`AdsApi`](../ads-api/) object which is used to create and manage ads.
```go
ads := api.Ads()
fmt.Println(ads) // AdsApi
```

### PluginsMgr
It returns the [`PluginsMgrApi`](../plugins-mgr-api/) object which is used to manage plugins.
```go
pluginsMgr := api.PluginsMgr()
fmt.Println(pluginsMgr) // PluginsMgrApi
```

### Network
It returns the [`NetworkApi`](../network-api/) object which is used to manage the network.
```go
network := api.Network()
fmt.Println(network) // NetworkApi
```

### DeviceHooks
It returns the [`DeviceHooksApi`](../device-hooks-api/) object which is used to manage device registration hooks.
```go
deviceHooks := api.DeviceHooks()
fmt.Println(deviceHooks) // DeviceHooksApi
```

### SessionsMgr
It returns the [`SessionsMgrApi`](../sessions-mgr-api/) object which is used to manage user sessions.
```go
sessionsMgr := api.SessionsMgr()
fmt.Println(sessionsMgr) // SessionsMgrApi
```

### Uci
It returns the [`UciApi`](../uci-api/) object which is a wrapper to [OpenWRT's UCI](https://openwrt.org/docs/guide-user/base-system/uci).
```go
uci := api.Uci()
fmt.Println(uci) // UciApi
```

### Themes
It returns the [`ThemesApi`](../themes-api/) object which is used to manage system UI themes.
```go
themes := api.Themes()
fmt.Println(themes) // ThemesApi
```
