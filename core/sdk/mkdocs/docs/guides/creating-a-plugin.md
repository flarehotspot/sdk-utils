
# Creating A Plugin

In this tutorial, we will create our first plugin. Open a terminal inside the devkit directory and type:

```sh
.\bin\flare.exe create-plugin
```

To start with, create a directory called `com.sample.plugin` under `plugins` directory of the sdk.
You can replace `com.sample.plugin` with your own package name.

```bash
cd ~/Documents/devkit-0.0.5
mkdir -p plugins/com.sample.plugin
```

Then, create a `plugin.json` file in the `com.sample.plugin` directory with the following content.

```json
{
  "package": "com.sample.plugin",
  "name": "Sample Plugin",
  "description": "This is a sample plugin",
  "version": "0.0.1",
  "author": "John Doe",
  "license": "MIT",
}
```

Then, create a `main.go` file in the `com.sample.plugin` directory with the following content.
```go
package main
import sdkplugin "github.com/flarehotspot/core/sdk/api/plugin"

func Init(api sdkplugin.PluginApi) {
    // Your plugin code goes here
}
```

Lastly, create a `go.mod` file in the `com.sample.plugin` directory with the following content. Replace `github.com/my-github-account/com.sample.plugin` with your own github account and package name.
```go
module github.com/my-github-account/com.sample.plugin

go 1.19
```