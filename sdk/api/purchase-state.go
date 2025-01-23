/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkapi

type PurchaseState struct {
	TotalPayment    float64 `json:"total_payment"`
	WalletDebit     float64 `json:"wallet_debit"`
	WalletEndingBal float64 `json:"wallet_ending_bal"`
	WalletRealBal   float64 `json:"wallet_real_bal"`
}
