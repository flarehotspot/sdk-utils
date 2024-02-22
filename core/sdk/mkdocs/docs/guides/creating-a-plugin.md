
# Creating a Plugin

In this tutorial, we will create our first plugin. In Windows, open a terminal inside the devkit directory. Then type:

```sh
.\bin\flare.exe create-plugin
```

If you are using Linux or Mac, type:
```sh
./bin/flare create-plugin
```

Follow the instructions in the command prompt and enter the necessary details for your plugin.
After that, your plugin will be created inside the `plugins` directory. Inside your plugin directory, you will find a `main.go` file. This file contains `Init` function which will be called when your plugin is loaded into the system. Below is the initial content of `main.go` file:

```go

package main

import (
	"net/http"

	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	sdkplugin "github.com/flarehotspot/core/sdk/api/plugin"
)

func main() {}

func Init(api sdkplugin.PluginApi) {
    // Your code goes here...
}
```

The `api` variable is an instance of the [PluginApi](../api/plugin-api.md), the root API of the Flare Hotspot SDK. Throughout the documentation, when you see `api`, it refers to this variable.
