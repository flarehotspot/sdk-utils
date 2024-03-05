# VueResponse

## VueResponse Methods

### Json

Used to send a JSON response to the client.
```go
data := map[string]string{
    "message": "Hello, World!",
}
api.Http().VueResponse().Json(w, data, http.StatusOK)
```

### FlashMsg

Used to send a pop up message to the client.

```go
api.Http().VueResponse().FlashMsg(w, "success", "Hello, World!")
```
