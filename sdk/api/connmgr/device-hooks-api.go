/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkconnmgr

import "net/http"

type ClientFinderFn func(r *http.Request, mac string, ip string, hostname string) (clnt ClientDevice, ok bool)
type ClientCreatedHookFn func(r *http.Request, clnt ClientDevice) error
type ClientChangedHookFn func(r *http.Request, current ClientDevice, old ClientDevice) error

type DeviceHooksApi interface {
	ClientFinderHook(...ClientFinderFn)
	ClientCreatedHook(...ClientCreatedHookFn)
	ClientChangedHook(...ClientChangedHookFn)
}
