# Form Submission

## Post Request Handler
We need a route and handler to handle the submitted form data.

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

## Form Component

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
            submit: function(){
                var self = this;
                var formData = {amount: self.amount};
                $flare.http.post('<% .Helpers.VueRoutePath "payment.received" %>', formData)
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
