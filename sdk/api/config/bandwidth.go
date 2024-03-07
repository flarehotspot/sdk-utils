/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkcfg

// BandwdCfg is the bandwidth configuration for a given interface. Each interface bandwidth is configured individually.
type BandwdCfg struct {
	// UseGlobal is true if the global bandwidth should be used.
	UseGlobal bool

	// GlobalDownMbits is the global download bandwidth in Mbits.
	GlobalDownMbits int

	// GlobalUpMbits is the global upload bandwidth in Mbits.
	GlobalUpMbits int

	// UserDownMbits is the per user download bandwidth in Mbits.
	UserDownMbits int

	// UserUpMbits is the per user upload bandwidth in Mbits.
	UserUpMbits int
}

// BandwidthCfgApi is used to get and set bandwidth configuration.
type BandwidthCfgApi interface {
	// Read returns the bandwidth configuration for a given interface.
	Get() (cfg BandwdCfg, ok bool)

	// SetConfig sets the bandwidth configuration for a given interface.
	// It needs application restart for the changes to take effect.
	Save(cfg BandwdCfg) error
}
