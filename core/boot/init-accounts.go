package boot

import "github.com/flarehotspot/flarehotspot/core/accounts"

func InitAccounts() {
	accounts.EnsureAdminAcct()
}
