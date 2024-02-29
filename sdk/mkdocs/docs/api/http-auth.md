# HttpAuth

The `HttpAuth` is used to authenticate and authorize admin users.

## HttpAuth Methods
The following are the available methods in `HttpAuth`.

### CurrentAcct
It returns the current admin user [Account](../accounts-api/#account-instance) instance from http request and an `error` if any. This method is only applicable on handlers registered on the [AdminRouter](../http-api/#admin-router).
```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    acct, err := api.Http().Auth().CurrentAcct(r)
    if err != nil {
        // handle error
    }
    fmt.Sprintf("Admin: %s", acct.Username) // Account
}
```

### Authenticate
It authenticates an account with a username and password. It returns an [Account](../accounts-api/#account-instance) instance and an `error` if any. This method is only applicable on handlers registered on the [PluginRouter](../http-api/#plugin-router), otherwise the request is blocked by the authentication middleware.
```go
// handler
func (r http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    username := r.PostFormValue("username")
    password := r.PostFormValue("password")
    acct, err := api.Http().Auth().Authenticate(username, password)
    if err != nil {
        // handle error
    }
    // proceed to api.Http().Auth().SignIn()
}
```

### SignIn
It signs in an account with an [Account](../accounts-api/#account-instance) instance by setting a cookie in the http response header. It returns an `error` if any. This method is only applicable on handlers registered on the [PluginRouter](../http-api/#plugin-router), otherwise the request is blocked by the authentication middleware.
```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    acct, err := api.Http().Auth().Authenticate("admin", "admin")
    if err != nil {
        // handle error
    }

    // set cookie header in the http response
    err = api.Http().Auth().SignIn(w, acct)
    if err != nil {
        // handle error
    }
    w.WriteHeader(http.StatusOK)
}
```

### SignOut
It signs out an [Account](../accounts-api/#account-instance) by removing the cookie from the http response header. It returns an `error` if any. This method works on any router.
```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    err := api.Http().Auth().SignOut(w)
    if err != nil {
        // handle error
    }
    w.WriteHeader(http.StatusOK)
}
```

