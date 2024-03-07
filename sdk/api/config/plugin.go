/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkcfg

// PluginCfgApi is used to read and write custom configuration for your plugin.
type PluginCfgApi interface {
	// Read reads the custom configuration of your plugin.
	// It is up to you to unmarshal the configuration into a struct.
	Get(v any) error

	// Writes writes the custom configuration of your plugin.
	// You have to marshal the configuration into a byte slice first.
	Save(v any) error
}
