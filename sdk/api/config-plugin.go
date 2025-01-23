/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkapi

type IPluginCfgApi interface {
	// Write a value to the plugin configuration file identified by key.
	Write(key string, value []byte) error

	// Read a value from the plugin configuration file identified by key.
	Read(key string) ([]byte, error)
}
