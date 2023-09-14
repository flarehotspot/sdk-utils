package config

// IPluginCfg is used to read and write custom configuration for your plugin.
type IPluginCfg interface {
	// Read reads the custom configuration of your plugin.
	// It is up to you to unmarshal the configuration into a struct.
	Read() ([]byte, error)

	// Writes writes the custom configuration of your plugin.
	// You have to marshal the configuration into a byte slice first.
	Write([]byte) error
}
