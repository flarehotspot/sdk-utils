# Creating a Link


## RouterLink
A router link is a vue component that's part of the official [vue-router](https://github.com/vuejs/vue-router) package. We can create a link to a [portal route](./basic-routing.md#portal-routes) or an [admin route](./basic-routing.md#admin-routes) by using the [HttpHelpers.VueRoutePath](../api/http-helpers.md#vueroutepath) method.

```html title="AnotherPage.vue"
<router-link :to='<% .Helpers.VueRoutePath "portal.welcome" "name" "Jhon" %>'>Go to welcome page</router-link>
```
This creates a link to the portal route named `portal.welcome` with a param `name` of value `Jhon`.

## Route Params
Route params can be passed to the [Helpers.VueRoutePath](../api/http-helpers.md#vueroutepath) as key-value pairs. For example you have a route path `/users/:user_id/posts/:post_id`, and the [name](./basic-routing.md#routename) of the path is `user.posts`, this is how you can create a link to that route with params:
```html
<router-link :to='<% .Helpers.VueRoutePath "user.posts" "user_id" "1" "post_id" "2" %>'> User posts </router-link>
```

