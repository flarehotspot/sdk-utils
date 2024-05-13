/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkcfg

// ConfigApi is used to access the configuration API.
type ConfigApi interface {

	// Get the application configuration api.
	Application() AppCfgApi

	// Get the bandwidth configuration api of a network interface.
	Bandwidth(ifname string) BandwidthCfgApi

	// Get the custom configuration guration api.
	Custom(key string) CustomCfgApi
}
