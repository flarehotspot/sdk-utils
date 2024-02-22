package sdkconnmgr

import "net/http"

type ClientFindHookFn func(w http.ResponseWriter, r *http.Request, mac string, ip string, hostname string) (ClientDevice, error)
type ClientCreatedHookFn func(w http.ResponseWriter, r *http.Request, clnt ClientDevice) error
type ClientChangedHookFn func(w http.ResponseWriter, r *http.Request, current ClientDevice, old ClientDevice) error
type ClientModifierHookFn func(w http.ResponseWriter, r *http.Request, current ClientDevice) (ClientDevice, error)

type DeviceHooksApi interface {
	ClientFindHook(ClientFindHookFn)
	ClientCreatedHook(ClientCreatedHookFn)
	ClientChangedHook(ClientChangedHookFn)
	ClientModifierHook(ClientModifierHookFn)
}
