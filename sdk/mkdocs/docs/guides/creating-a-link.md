# Creating a Link

```html title="AnotherPage.vue"
<router-link :to='<% .Helpers.VueRoutePath "portal.welcome" "name" "Jhon" %>'>Go to welcome page</router-link>
```
This creates a link to the portal route named `portal.welcome` with a param `name` of value `Jhon`.
