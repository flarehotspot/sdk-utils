
# IHttpFormsApi

The `IHttpFormsApi` is used to build HTML forms. It is responsible for rendering, validating and parsing the HTML form and its values.


## IHttpFormsApi methods

Below are the methods available in `IHttpFormsApi`.

### RegisterForms

It registers one or more [HttpForm](#httpform) into the plugin.

```go
form := sdkapi.HttpForm{
    Name: "my-form", // name of the form
    // rest of the form definition...
}
formsAPI := api.Http().Forms()
if err := formsAPI.RegisterForms(form); err != nil {
    // handle error
}
```

### GetForm

It returns an instance of a registered [IHttpForm](#ihttpform).

```go
formsAPI := api.Http().Forms()
form, ok := formsAPI.GetForm("my-form")
if !ok {
    // handle error
}
```

## The `HttpForm` struct {#httpform}

The `HttpForm` struct defines the HTML form sections, fields, default input values, and validation rules.
It is composed of one or more sections. Each section contains various types of fields which include [text](#text-field), [decimal](#decimal-field), [integer](#integer-field), [boolean](#boolean-field), [list](#list-field) and [multi-field](#multi-field).

Below is an example of `HttpForm` definition:

```go
pluginConfigAPI := api.Config().Plugin()
formsAPI := api.Http().Forms()

sections := []sdkapi.FormSection{
    {
        {
            Name: "General Configuration",
            Fields: []sdkapi.IFormField{
                // sdkapi.FormBooleanField,
                // sdkapi.FormDecimalField,
                // sdkapi.FormIntegerField,
                // sdkapi.FormListField,
                // sdkapi.FormMultiField,
                sdkapi.FormTextField{
                    Name:  "Banner Text",
                    Label: "Banner Text",
                    ValueFn: func() string {
                        b, err := pluginConfigAPI.Read("banner_text")
                        if err != nil {
                            return "This is the default banner text!"
                        }
                        return string(b)
                    },
                },
            },
        },
    },
}

form := sdkapi.HttpForm{
    Name: "my-form", // name of the form
    CallbackRoute: "settings:save", // route to handle form submission
    SubmitLabel: "Submit", // submit button text
    Sections: sections,
}

formsAPI.RegisterForms(form)
```

## Form Fields {#form-fields}

Below are the available fields that can be used in the `HttpForm` definition.

### Boolean Field

TODO: Add description

### Decimal Field

TODO: Add description

### Integer Field

### List Field

### Multi Field

### Text Field

## IHttpForm methods {#ihttpform}

### GetTemplate

It returns the templ component of the form which can be used to render to the view.

```go
// handler
func handler(w http.ResponseWriter, r *http.Request) {
    formsAPI := api.Http().Forms()
    form, _ := formsAPI.GetForm("my-form")
    formHTML := form.GetTemplate(r)
    // render formHTML to the view
}
```

### GetSections

Returns a slice of [FormSection](#formsection) in the form.

```go
formsAPI := api.Http().Forms()
form, _ := formsAPI.GetForm("my-form")
sectoins := form.GetSections()
```

### GetStringValue

Returns the string value of a field in the form.

```go
val, err := form.GetStringValue(section, "Banner Text")
```

### GetStringValues

Returns a slice of strings for list or multi fields.

```go
vals, err := form.GetStringValues(section, "List Field")
```

### GetIntValue

Returns the integer value of a field in the form.

```go
val, err := form.GetIntValue(section, "Integer Field")
```

### GetIntValues

Returns a slice of integers for list or multi fields.

```go
vals, err := form.GetIntValues(section, "Integer List")
```

### GetFloatValue

Returns the float value of a decimal field in the form.

```go
val, err := form.GetFloatValue(section, "Decimal Field")
```

### GetFloatValues

Returns a slice of floats for list or multi fields.

```go
vals, err := form.GetFloatValues(section, "Decimal List")
```

### GetBoolValue

Returns the boolean value of a field in the form.
```go
val, err := form.GetBoolValue(section, "Boolean Field")
```

### GetBoolValues

Returns a slice of booleans for list or multi fields.

```go
vals, err := form.GetBoolValues(section, "Boolean List")
```

### GetMultiField

Returns a [IFormMultiField](#imultifield) instance of a multi field in the form.

```go
mf, err := form.GetMultiField(section, "Multi Field")
```

## FormSection struct {#formsection}

A `FormSection` is a collection of fields in a form. It is composed of a name and a slice of [Form Fields](#form-fields).

## IFormMultiField {#imultifield}
