/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkcfg

// ConfigApi is used to access the configuration API.
type ConfigApi interface {

	// Get the plugin configuration guration api.
	Plugin() PluginCfg

	// Get the application configuration api.
	Application() ApplicationCfg

	// Get the database configuration api.
	Database() DbCfg

	// Get the http session rates configuration api.
	SessionRates() SessionRatesCfg

	// Get the sessions configuration api.
	Sessions() SessionLimitsCfg

	// Get the bandwidth configuration api.
	Bandwidth() BandwidthCfg
}
