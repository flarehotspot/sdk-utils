
# Creating a Plugin

## The `create-plugin` command

To create a new plugin, open a terminal and navigate inside the devkit directory.

If you are using Windows `CMD` or `PowerShell`, type:
```cmd title="PowerShell"
.\scripts\flare.bat create-plugin
```

If you are using Linux/Mac/WSL, type:
```sh title="Terminal"
./scripts/flare.sh create-plugin
```

Follow the instructions in the command prompt and enter the necessary details for your plugin. Below are the needed details for your plugin:

Package Name
: This is the primary identifier of your plugin. It should be unique and follow reverse domain naming convention, e.g `com.mydomain.myplugin`. The package name should be in lowercase and should not contain any special characters or spaces except period, underscore and hyphen (`.`, `_`, `-`).

Plugin Name
: This is the name of your plugin, e.g. "System Monitor".

Description
: This is a brief description of your plugin. It should describe the purpose of your plugin.

## The main.go file

After that, your plugin will be created inside the `plugins/local/[your-plugin-package]` directory (`plugins/local/com.mydomain.myplugin` in this example). Inside your plugin directory, you will find a `main.go` file.

![main.go file](./img/main-go-location.png)

This file contains `Init` function which will be called when your plugin gets loaded into the system. Below is the initial content of `main.go` file:

```go title="main.go"

package main

import (
	"net/http"

	sdkapi "sdk/api"
)

func main() {}

func Init(api sdkapi.PluginApi) {
    // Rest of the code...
}
```

!!! note
    The `api` variable is an instance of the [IPluginApi](../api/plugin-api.md), the root API of the Flare Hotspot SDK. Throughout the documentation, when you see the variable `api`, it refers to [IPluginApi](../api/plugin-api.md).

## Troubleshooting

For linux users, you must change the file permissions to fix errors in your code editor:
```sh title="Terminal"
sudo chown -R $USER .
```

For MacOS users, if you encouter `Too many open files in system` error, you can fix this by cleaning the Go build cache and fixing the file permissions:

```sh title="Terminal"
go clean -cache
sudo chown -R $USER .
```
