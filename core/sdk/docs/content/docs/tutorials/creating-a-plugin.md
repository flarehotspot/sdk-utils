+++
title = "Creating A Plugin"
description = "In this tutorial, we will create our first plugin by creating a new directory under plugins directory of the sdk."
date = 2021-05-01T08:00:00+00:00
updated = 2021-05-01T08:00:00+00:00
draft = false
weight = 110
sort_by = "weight"
template = "docs/page.html"

+++

# Creating A Plugin

In this tutorial, we will create our first plugin by creating a new directory under `plugins` directory of the sdk.
The directory name must match the `package` field in your [plugin.json](../api/plugin-json) file.

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
