# PluginApi

The `PluginApi` provides access to methods used to manipulate system accounts, network devices, theme configuration, user sessions and payment system. Each plugin is provided with an instance of `PluginApi`, the root interface of our SDK.

When the plugin is first loaded into the system, the system looks for the `Init` function of the plugin's `main` package. The `PluginApi` object is then passed to the init function. From here, you can start configuring the routes and components of your plugin. An example of a plugin's init function:

```go
// file: plugins/com.mydomain.myplugin/main.go

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

---

The following are the available methods in `PluginApi`.

## Name
It returns the `name` field defined in `plugin.json`.

```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    name := api.Name()
    fmt.Println(name) // "My Plugin"
}
```

## Pkg
It returns the `package` field defined in [plugin.json](../plugin-json/).
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    pkg := api.Pkg()
    fmt.Println(pkg) // "com.mydomain.myplugin"
}
```

## Version
It returns the `version` field defined in [plugin.json](../plugin-json/).
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    version := api.Version()
    fmt.Println(version) // "1.0.0"
}
```

## Description
It returns the `description` field defined in [plugin.json](../plugin-json/).
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    description := api.Description()
    fmt.Println(description) // "My plugin description"
}
```

## Dir
It returns the absolute path of the plugin's installtion directory.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    dir := api.Dir()
    fmt.Println(dir) // "/path/to/com.mydomain.myplugin"
}
```

## Translate
It is a utility function used to convert a key into a translated string. Example usage:
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    msg := api.Translate("info", "payment_received", "amount", 1.00)
    fmt.Println(msg) // "Payment received USD 1.0.0"
}
```

In this example, given that the application language config is set to `en`, the system will look for the file `com.mydomain.myplugin/resources/translations/en/info/payment_received.txt`. If the file is found, the system will use the contents of the file as the translation template.

We also want the to replace the amount with the actual amount. We can do that by passing the amount param as key-value pairs (e.g "amount", 1.00) to the `Translate` method. Translation params are represented as `<% .param %>` (with a dot prefix) in the translation file. Therefore the content of `payment_received.txt` should be:
```go
Payment received: USD <% .amount %>
```

Internally, the param pairs are converted into a `map[any]any`.

## Resource
It returns the absolute path of the file under the plugin's resource directory.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    resource := api.Resource("/my-resource.txt")
    fmt.Println(resource) // "/path/to/com.mydomain.myplugin/resources/my-resource.txt"
}
```

## SqlDb
It returns [`*sql.DB`](http://go-database-sql.org/overview.html) instance which is used to query, insert, update and delete database entities.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    db := api.SqlDb()
    fmt.Println(db) // *sql.DB
}
```

## Acct
It returns the [AccountsApi](../accounts-api/) object which is used to access and modify the system admin accounts.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    acct := api.Acct()
    fmt.Println(acct) // AccountsApi
}
```

## Http
It returns the [`HttpApi`](../http-api/) object which is used to configure routes and serve HTTP requests.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    http := api.Http()
    fmt.Println(http) // HttpApi
}
```

## Config
It returns the [`ConfigApi`](../config-api/) object which is used to access and modify the system configuration.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    config := api.Config()
    fmt.Println(config) // ConfigApi
}
```

## Payments
It return the [`PaymentsApi`](../payments-api/) object which is used to create payment options or create system transactions.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    payments := api.Payments()
    fmt.Println(payments) // PaymentsApi
}
```

## InAppPurchases
It returns the [`InAppPurchasesApi`](../in-app-purchases-api/) object which is used to create and manage in-app purchases.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    inAppPurchases := api.InAppPurchases()
    fmt.Println(inAppPurchases) // InAppPurchasesApi
}
```

## Ads
It returns the [`AdsApi`](../ads-api/) object which is used to create and manage ads.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    ads := api.Ads()
    fmt.Println(ads) // AdsApi
}
```

## PluginsMgr
It returns the [`PluginsMgrApi`](../plugins-mgr-api/) object which is used to manage plugins.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    pluginsMgr := api.PluginsMgr()
    fmt.Println(pluginsMgr) // PluginsMgrApi
}
```

## Network
It returns the [`NetworkApi`](../network-api/) object which is used to manage the network.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    network := api.Network()
    fmt.Println(network) // NetworkApi
}
```

## DeviceHooks
It returns the [`DeviceHooksApi`](../device-hooks-api/) object which is used to manage device registration hooks.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    deviceHooks := api.DeviceHooks()
    fmt.Println(deviceHooks) // DeviceHooksApi
}
```

## SessionsMgr
It returns the [`SessionsMgrApi`](../sessions-mgr-api/) object which is used to manage user sessions.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    sessionsMgr := api.SessionsMgr()
    fmt.Println(sessionsMgr) // SessionsMgrApi
}
```

## Uci
It returns the [`UciApi`](../uci-api/) object which is a wrapper to [OpenWRT's UCI](https://openwrt.org/docs/guide-user/base-system/uci).
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    uci := api.Uci()
    fmt.Println(uci) // UciApi
}
```

## Themes
It returns the [`ThemesApi`](../themes-api/) object which is used to manage system UI themes.
```go
package main
// truncated code...
func Init(api sdkplugin.PluginApi) {
    themes := api.Themes()
    fmt.Println(themes) // ThemesApi
}
```
