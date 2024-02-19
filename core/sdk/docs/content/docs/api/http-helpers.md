+++
title = "Http Helpers"
description = "Utilty functions for handling http requests and responses."
date = 2021-05-01T08:00:00+00:00
updated = 2021-05-01T08:00:00+00:00
draft = false
weight = 20
sort_by = "weight"
template = "docs/page.html"

[extra]
lead = "Admin authentication and authorization API."
toc = true
top = false
+++

# HttpHelpers

## Overview
The `HttpHelpers` consist of utility functions for the [views](../resources/#views) and [vue component](../resources/#vue-components) files. It can also be used in your handlers using [HttpApi.Helpers](../http-api/#helpers)

# Methods
First, get an instance of the `HttpHelpers` from the [HttpApi](../http-api/#helpers):
```go
package main
// imports...
func Init(api sdkplugin.PluginApi) {
    httpApi := api.Http()
    helpers := httpApi.Helpers()
}
```
