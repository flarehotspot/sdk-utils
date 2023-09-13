package boot

import "github.com/flarehotspot/core/accounts"

func InitAccounts() {
	accounts.EnsureAdminAcct()
}
