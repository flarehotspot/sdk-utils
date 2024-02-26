/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkcfg

// SessionRate is used to compute the session's time and data.
type SessionRate struct {
	// Uuid a custom identifier key for this rate config.
	Uuid string

	// Network is the IP CIDR for which this rate is applicable.
	Network string

	// Amount is the cost for internet time and/or data.
	Amount float64

	// TimeMins is the equivalent time in minutes for the amount.
	TimeMins uint

	// DataMbytes is the equivalent data in megabytes for the amount.
	DataMbytes uint
}

// SessionRatesCfg is the configuration for internet connection rates.
type SessionRatesCfg interface {

	// Returns all rates.
	All() (rates []SessionRate, err error)

	// Returns all rates for a given network.
	AllByNet(lanIP string) ([]SessionRate, error)

	// Updates or creates a new rate if its Uuid does not exist.
	Save(rate SessionRate) error

	// Saves the new rates configuration.
	Write(rates []SessionRate) ([]SessionRate, error)

	// Returns the session's time and data based on the amount and session type.
	ComputeSession(clientIP string, amount float64, t uint8) (result SessionResult, err error)
}

// SessionResult is the result of the computation base on the amount and session type.
type SessionResult struct {
	TimeMins   uint
	DataMbytes uint
}
