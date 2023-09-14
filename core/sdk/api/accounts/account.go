package accounts

// IAccount represents a system account.
type IAccount interface {

  // Returns the username for this account.
	Username() string

  // Returns true if the password is correct.
	Auth(password string) bool

	// Check if this account is an admin.
	IsAdmin() bool

	// Get the permissions for this account.
	Permissions() []string

	// Update this account.
	Update(username string, password string, permissions []string) error

	// Delete this account.
	Delete() error

	// Emit events to the browser for this account.
  // Events will be propagated to the client's browser via server-sent events.
	Emit(event string, data interface{})
}
