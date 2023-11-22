package account_manager_proxy

import "github.com/netbirdio/netbird/management/server"

type AccountManagerProxy struct {
	server.AccountManager
}

func NewAccountManagerProxy(accountManager server.AccountManager) *AccountManagerProxy {
	return &AccountManagerProxy{
		accountManager,
	}
}
