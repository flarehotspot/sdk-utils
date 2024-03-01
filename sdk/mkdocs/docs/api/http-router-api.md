# HttpRouterApi
The `HttpRouterApi` is the backend for http routing including. The [VueRouterApi](./vue-router-api.md) uses the `HttpRouterApi` to generate the routes for the frontend. Each plugin are provided with a `HttpRouterApi` instance to generate their own routes.

## HttpRouterApi Methods
Below are the available methods in `HttpRouterApi`:

### PluginRouter
This method returns the [plugin router instance](./http-router-instance.md) for the plugin routes. Routes generated from the plugin router are prefixed with `/plugin` and are accessible to all users.

### AdminRouter
This method returns the [admin router instance](./http-router-instance.md) for the admin routes. Routes generated from the admin router are prefixed with `/admin` and are only accessible to authenticated user [accounts](./accounts-api.md#account-instance).

### MuxRouteName

### UrlForMuxRoute

### UrlForRoute

### UrlForPkgRoute
