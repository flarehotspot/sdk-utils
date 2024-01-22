package sdkplugin

import (
	"github.com/flarehotspot/core/sdk/utils/translate"
)

type IPluginUtils interface {
	// Translates the given message key to the current language.
	// This is the same and is identical to the view helper's "Translate()" method.
	// For example, if the current language is "en", then the following code:
	//  api.Translate(translate.Error, "some-key")
	// will look for the file "/resources/translations/en/error/some-key.txt" under the plugin's root directory
	// and displays the text inside that file.
	Translate(t sdktrans.MsgType, msgk string) string

	// Returns the absolute path to the given file in /resources folder of your plugin.
	// For example, if you have the following code:
	//  api.Resource("some-file.txt")
	// then it will return the absolute path to the file "/resources/some-file.txt" under the plugin's root directory.
	Resource(f string) (path string)
}
