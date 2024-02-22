package sdkacct

// AccountsApi is used to manage accounts.
type AccountsApi interface {

	// Create a new system account. The list of available permissions
	// can be obtained from IAcctApi.Permissions().
	Create(username string, pass string, perms []string) (Account, error)

	// Find an account by username.
	Find(username string) (Account, error)

	// Get all accounts, admin and non-admin.
	GetAll() ([]Account, error)

	// Get all admin accounts.
	GetAdmins() ([]Account, error)

	// Add a new type of permission.
	NewPerm(name string, desc string) error

	// Retrieve all permissions.
	GetPerms() map[string]string

	// Retrieve a permission description.
	PermDesc(perm string) (desc string)
}
