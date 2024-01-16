package plugins

import (
	"net/http"

	"github.com/flarehotspot/core/accounts"
	acct "github.com/flarehotspot/core/sdk/api/accounts"
	"github.com/flarehotspot/core/web/helpers"
)

type AccountsApi struct {
	api *PluginApi
}

func NewAcctApi(api *PluginApi) *AccountsApi {
	return &AccountsApi{api}
}

func (self *AccountsApi) CurrentAdmin(r *http.Request) (acct.IAccount, error) {
	return helpers.CurrentAdmin(r)
}

func (self *AccountsApi) Create(uname string, pass string, perms []string) (acct.IAccount, error) {
	return accounts.Create(uname, pass, perms)
}

func (self *AccountsApi) Update(oldname string, uname string, pass string, perms []string) (acct.IAccount, error) {
	return accounts.Update(oldname, uname, pass, perms)
}

func (self *AccountsApi) Delete(uname string) error {
	return accounts.Delete(uname)
}

func (self *AccountsApi) AllAccounts() ([]acct.IAccount, error) {
	accts, err := accounts.All()
	if err != nil {
		return nil, err
	}

	accounts := []acct.IAccount{}
	for _, a := range accts {
		accounts = append(accounts, a)
	}

	return accounts, nil
}

func (self *AccountsApi) AllAdmin() ([]acct.IAccount, error) {
	accts, err := accounts.All()
	if err != nil {
		return nil, err
	}

	admins := []acct.IAccount{}
	for _, acct := range accts {
		if acct.IsAdmin() {
			admins = append(admins, acct)
		}
	}

	return admins, nil
}

func (self *AccountsApi) Find(username string) (acct.IAccount, error) {
	return accounts.Find(username)
}

func (self *AccountsApi) NewPerm(name string, desc string) error {
	return accounts.NewPerm(name, desc)
}

func (self *AccountsApi) Permissions() map[string]string {
	return accounts.Permissions()
}

func (self *AccountsApi) PermDesc(name string) (desc string) {
	return accounts.PermDesc(name)
}

func (self *AccountsApi) HasAllPerms(acct acct.IAccount, perms ...string) bool {
	return accounts.HasAllPerms(acct, perms...)
}

func (self *AccountsApi) HasAnyPerm(acct acct.IAccount, perms ...string) bool {
	return accounts.HasAnyPerm(acct, perms...)
}
