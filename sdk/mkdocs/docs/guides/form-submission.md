# Form Submission

In this tutorial, we will create a form component and submit the form data to the server using the [$flare.http.post](../api/flare-variable.md#flare-http-post) method.
The [$flare](../api/flare-variable.md) variable is a global variable in the browser that contains helper functions to work with the Flare API.

## 1. Request Handler {#request-handler}
We need a route and handler to handle the submitted form data. Below is an example of a route and handler to handle the form submission.

```go
api.Http().HttpRouter().PluginRouter().Post("/payments/receieved", func (w http.ResponseWriter, r *http.Request) {
  var data struct {
    Amount int `json:"amount"`
  }
  if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  fmt.Fprintf(w, "Payment Received: %d", data.Amount)

}).Name("payment.received")
```

In this example, we are using the `Post` method from [PluginRouter](../api/http-router-api.md#post) to handle the form submission. The first argument is the route path, and the second argument is the handler function. The handler function accepts two arguments, the first argument is the `http.ResponseWriter`, and the second argument is the `*http.Request`.

Then we are decoding the form data using the `json.NewDecoder(r.Body).Decode(&data)` method. The `json.NewDecoder(r.Body).Decode(&data)` method decodes the form data from the request body and stores it in the `data` variable.

Lastly, the `Name` method is used to give the route a name.

## 2. Form Component {#form-component}

To submit a form, we need to prevent the default form submission by using the `@click.prevent` directive on the form tag.
Then, we will use the `$flare.http.post` method to send the form data to the server.
The `$flare.http.post` method accepts two arguments, the first argument is the URL to send the form data to, and the second argument is the form data.

```html title="resources/components/Form.vue"
<template>
  <form @click.prevent="submit">
    <label>Amount:</label>
    <input type="text" v-model="amount">
    <button type="submit">Submit</button>
  </form>
</template>

<script>
define(function(){
  return {
    props: ['flareView'],
    template: template,
    data: function() {
      return {amount: 0};
    },
    methods: {
      submit: function() {
        var self = this;
        var formData = {amount: self.amount};

        $flare.http.post('<% .Helpers.UrlForRoute "payment.received" %>', formData)
          .then(function(response){
            console.log(response);
          })
          .catch(function(error){
            console.log(error);
          });
      }
    }
  };
});
</script>
```
