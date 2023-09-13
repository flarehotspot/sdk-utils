package connmgr

import "net/http"

type ClientFindHookFn func(w http.ResponseWriter, r *http.Request, mac string, ip string, hostname string) (IClientDevice, error)
type ClientCreatedHookFn func(w http.ResponseWriter, r *http.Request, clnt IClientDevice) error
type ClientChangedHookFn func(w http.ResponseWriter, r *http.Request, current IClientDevice, old IClientDevice) error
type ClientModifierHookFn func(w http.ResponseWriter, r *http.Request, current IClientDevice) (IClientDevice, error)

type IClientRegister interface {
  CurrentClient(r *http.Request) (IClientDevice, error)
	ClientFindHook(ClientFindHookFn)
	ClientCreatedHook(ClientCreatedHookFn)
	ClientChangedHook(ClientChangedHookFn)
	ClientModifierHook(ClientModifierHookFn)
}
