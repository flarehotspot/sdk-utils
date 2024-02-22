# HttpAuth

The `HttpAuth` is used to authenticate and authorize admin users.

# Methods
First, get an instance of the `HttpAuth` from the [HttpApi](../http-api/#auth):
```go
package main
// imports...
func Init(api sdkplugin.PluginApi) {
    httpApi := api.Http()
    authApi := httpApi.Auth()
}
```

The following are the available methods in `HttpAuth`.

## CurrentAcct
It returns the current admin user [Account](../accounts-api/#account-instance) instance from http request and an `error`. This method is only applicable on handlers registered on the [AdminRouter](../http-api/#admin-router).
```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    acct, err := authApi.CurrentAcct(r)
    if err != nil {
        // handle error
    }
    fmt.Sprintf("Admin: %s", acct.Username) // Account
}
```

## Authenticate
It authenticates an admin user with a username and password. It returns an [Account](../accounts-api/#account-instance) instance and an `error`. This method is only applicable on handlers registered on the [PluginRouter](../http-api/#plugin-router), otherwise the request is blocked by the authentication middleware.
```go
// handler
func (r http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    username := r.PostFormValue("username")
    password := r.PostFormValue("password")
    acct, err := authApi.Authenticate(username, password)
    if err != nil {
        // handle error
    }
    // proceed to authApi.SignIn()
}
```

## SignIn
It signs in an admin user with an [Account](../accounts-api/#account-instance) instance by setting a cookie in the http response header. It returns an `error`. This method is only applicable on handlers registered on the [PluginRouter](../http-api/#plugin-router), otherwise the request is blocked by the authentication middleware.
```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    acct, err := authApi.Authenticate("admin", "admin")
    if err != nil {
        // handle error
    }

    // set cookie header in the http response
    err = authApi.SignIn(w, acct)
    if err != nil {
        // handle error
    }
    w.WriteHeader(http.StatusOK)
}
```

## SignOut
It signs out an admin user by removing the cookie from the http response header. It returns an `error`. This method works on any router.
```go
// handler
func (w http.ResponseWriter, r *http.Request) {
    err := authApi.SignOut(w)
    if err != nil {
        // handle error
    }
    w.WriteHeader(http.StatusOK)
}
```

