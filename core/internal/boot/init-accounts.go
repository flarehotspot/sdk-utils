package boot

import "github.com/flarehotspot/core/internal/accounts"

func InitAccounts() {
	accounts.EnsureAdminAcct()
}
