package sdkplugin

import ()

type PluginUtils interface {

	// Translate a message to the user's language.
	Translate(t string, msgk string, pairs ...string) string

	// Returns the absolute path to the given file in /resources folder of your plugin.
	// For example, if you have the following code:
	//  api.Utils().Resource("some-file.txt")
	// then it will return the absolute path to the file "[plugin_root_dir]/resources/some-file.txt" under the plugin's root directory.
	Resource(f string) (path string)
}
