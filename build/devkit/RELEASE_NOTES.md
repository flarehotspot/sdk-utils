**Release Notes:**

## 1. Devkit auto-recompile plugin on file change

The server will now automatically recompile the plugin when a go file is changed in the plugin directory.
You no longer have to restart the docker container to see the changes.

## 2. Custom plugin config API change

The custom plugin config method `api.Config().Plugin("key")` is now replaced with `api.Config().Custom("key")`.

To retrieve the custom config data:

```go
// get custom config
config := struct {
    // .. struct fields
}

myConfig, err := api.Config().Custom("my_key").Get(&config)
```

To save custom config data:

```go
// get custom config
config := map[string]any{
    // .. config data
}

err := api.Config().Custom("my_key").Save(&config)
```
