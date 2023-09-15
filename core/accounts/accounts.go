package accounts

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/flarehotspot/core/sdk/api/accounts"
	"github.com/flarehotspot/core/sdk/libs/yaml-3"
	"github.com/flarehotspot/core/sdk/utils/fs"
	"github.com/flarehotspot/core/sdk/utils/paths"
	"github.com/flarehotspot/core/sdk/utils/translate"
)

const (
	AdminUsername = "admin"
	AdminPassword = "admin"
	PermMngUsers  = "manage_users"
)

var (
	perms        sync.Map
	DefaultPerms = []string{PermMngUsers}
)

func init() {
	perms = sync.Map{}
	perms.Store(PermMngUsers, translate.Core(translate.Label, "perm_manage_users"))
}

func DefaultAdminAcct() Account {
	var acct Account
	f := filepath.Join(paths.DefaultsDir, "admin.yml")

	perms := []string{}
	for _, p := range DefaultPerms {
		perms = append(perms, p)
	}

	defAcct := Account{
		Uname:  AdminUsername,
		Passwd: AdminPassword,
		Perms:  perms,
	}

	bytes, err := os.ReadFile(f)
	if err != nil {
		return defAcct
	}

	err = yaml.Unmarshal(bytes, &acct)
	if err != nil {
		return defAcct
	}

	return acct
}

func EnsureAdminAcct() error {
	f := FilepathForUser(AdminUsername)
	if !fs.Exists(f) {
		acct := DefaultAdminAcct()
		content, err := yaml.Marshal(acct)
		if err != nil {
			return err
		}
		err = os.WriteFile(f, content, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func All() (accounts []*Account, err error) {
	files, err := fs.LsFiles(AcctDir, false)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}

		var acct Account
		err = yaml.Unmarshal(b, &acct)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, &acct)
	}

	return accounts, err
}

func AllAdmins() ([]*Account, error) {
	accts, err := All()
	if err != nil {
		return nil, err
	}

	admins := []*Account{}
	for _, acct := range accts {
		if acct.IsAdmin() {
			admins = append(admins, acct)
		}
	}

	return admins, nil
}

func Find(username string) (*Account, error) {
	var acct Account
	derr := errors.New(translate.Core(translate.Error, "account_not_exist"))
	f := FilepathForUser(username)
	b, err := os.ReadFile(f)
	if err != nil {
		return nil, derr
	}
	err = yaml.Unmarshal(b, &acct)
	if err != nil {
		return &acct, derr
	}
	return &acct, nil
}

func Create(uname string, passwd string, perms []string) (*Account, error) {
	acct := Account{
		Uname:  uname,
		Passwd: passwd,
		Perms:  perms,
	}

	b, err := yaml.Marshal(&acct)
	if err != nil {
		return nil, err
	}

	f := FilepathForUser(uname)
	if fs.Exists(f) {
		return nil, fmt.Errorf("Account with username \"%s\" already exists", uname)
	}

	err = os.WriteFile(f, b, 0644)
	if err != nil {
		return nil, err
	}

	return &acct, nil
}

func Update(prevName string, newName string, pass string, perms []string) (*Account, error) {
	prevAcct, err := Find(prevName)
	if err != nil {
		return nil, err
	}

	if prevAcct.Uname == AdminUsername && newName != AdminUsername {
		return nil, errors.New("Renaming the super admin account is not allowed.")
	}

	if pass == "" {
		pass = prevAcct.Passwd
	}

	if len(perms) == 0 {
		perms = prevAcct.Perms
	}

	acct := &Account{
		Uname:  newName,
		Passwd: pass,
		Perms:  perms,
	}

	if acct.Uname == AdminUsername && !HasAllPerms(acct, PermMngUsers) {
		return nil, errors.New("Super admin account must have manage users permission.")
	}

	b, err := yaml.Marshal(&acct)
	if err != nil {
		return nil, err
	}

	f := FilepathForUser(newName)
	err = os.WriteFile(f, b, 0644)
	if err != nil {
		return nil, err
	}

	if prevName != newName {
		f = FilepathForUser(prevName)
		err = os.Remove(f)
		return acct, err
	}

	return acct, nil
}

func Delete(uname string) error {
	if uname == AdminUsername {
		return fmt.Errorf("Deleting the super admin account is not allowed.")
	}

	files, err := fs.LsFiles(AcctDir, false)
	if err != nil {
		return err
	}

	acct, err := Find(uname)
	if err != nil {
		return err
	}

	if len(files) < 2 && acct.Uname == uname {
		return errors.New("Can't delete last super admin user.")
	}

	return os.Remove(FilepathForUser(uname))
}

func FilepathForUser(uname string) string {
	return filepath.Join(AcctDir, uname+".yml")
}

// Permissions returns all permissions from perms.SyncMap as map[string]string
func Permissions() map[string]string {
	permsMap := map[string]string{}
	perms.Range(func(key, value interface{}) bool {
		permsMap[key.(string)] = value.(string)
		return true
	})
	return permsMap
}

// PermDesc returns description string of permission name
func PermDesc(perm string) string {
	desc, ok := perms.Load(perm)
	if !ok {
		return ""
	}
	return desc.(string)
}

// Check if account has all permissions
func HasAllPerms(acct accounts.IAccount, perms ...string) bool {
	count := 0
	for _, perm := range perms {
		for _, acctPerm := range acct.Permissions() {
			if perm == acctPerm {
				count++
			}
		}
	}

	return count == len(perms)
}

// Check if account has any of the permissions
func HasAnyPerm(acct accounts.IAccount, perms ...string) bool {
	for _, perm := range perms {
		for _, acctPerm := range acct.Permissions() {
			if perm == acctPerm {
				return true
			}
		}
	}
	return false
}

// Add new permission to perms sync.Map with name and description params
func NewPerm(name string, description string) error {
	_, ok := perms.Load(name)
	if ok {
		return errors.New("Permission with name \"" + name + "\" already exists")
	}

	perms.Store(name, description)
	return nil
}
