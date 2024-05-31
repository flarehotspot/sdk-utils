## Rename sdk module to only `sdk`

The code below shows the difference between the old and new import path for the sdk module:

```go
// Before:
// import sdkplugin "github.com/flarehotspot/sdk/api/plugin"

// Now:
import sdkplugin "sdk/api/plugin"
```
