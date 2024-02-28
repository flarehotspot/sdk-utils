# Translations

To translate texts, we will use the [Translate](../api/plugin-api.md#translate) method from [PluginApi](../api/plugin-api.md).

The `Translate` method receives a message type and message key string and returns the translated text. The method will look for the translation in the `resources/translations/[lang]/[type]/[file].txt` file in your plugin.

The `[lang]` placeholder is the language code set in the [application config](../api/config-api.md#application), e.g. `en` for English.

The `[type]` placeholder is the type of the translation message, e.g. `label` for labels and button texts. Other types are `info` and `error`.

The `[file]` placeholder is the message key of the translation message.

For example, to translate message key "save" to the target language, we can use the following code:
```go
saveText := api.Translate("label", "save")
```

In this example, the `Translate` method will look for the file `resources/translations/en/label/save.txt`. The contents of the file will be used as the template for the translated text. For more advanced translations, see [PluginApi.Translate](../api/plugin-api.md#translate) method documentation.
