/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkpayments

import (
	connmgr "sdk/api/connmgr"
)

// PaymentProvider represents a payment provider.
// A payment provider can have many payment options.
type PaymentProvider interface {

	// Returns name of the payment provider.
	Name() string

	// Returns a list of available payment options.
	PaymentOpts(clnt connmgr.ClientDevice) []PaymentOpt
}
