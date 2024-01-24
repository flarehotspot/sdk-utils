package sdkacct

// IAcctApi is used to manage accounts.
type IAcctApi interface {

	// Create a new system account. The list of available permissions
	// can be obtained from IAcctApi.Permissions().
	Create(username string, pass string, perms []string) (IAccount, error)

	// Get all accounts, admin and non-admin.
	AllAccounts() ([]IAccount, error)

	// Get all admin accounts.
	AllAdmin() ([]IAccount, error)

	// Find an account by username.
	Find(username string) (IAccount, error)

	// Update an existing account.
	Update(oldusername string, username string, pass string, perms []string) (IAccount, error)

	// Delete an account by username.
	Delete(username string) error

	// Add a new type of permission.
	NewPerm(name string, desc string) error

	// Retrieve all permissions.
	Permissions() map[string]string

	// Retrieve a permission description.
	PermDesc(perm string) (desc string)

	// Check if account has all of the specified permissions.
	HasAllPerms(acct IAccount, perms ...string) bool

	// Check if account has any of the specified permissions.
	HasAnyPerm(acct IAccount, perms ...string) bool
}
