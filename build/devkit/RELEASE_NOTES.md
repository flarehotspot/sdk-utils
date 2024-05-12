**Release Notes:**

## 1. Custom plugin config API change

Custom configurations can now be get and saved using new syntax:

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
