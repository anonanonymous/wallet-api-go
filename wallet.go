package walletapi

// Wallet - holds wallet configuration
type Wallet struct {
	DaemonHost string
	DaemonPort int
	Filename   string
	Password   string
}

// OpenWallet - opens an existing wallet
func (wAPI WalletAPI) OpenWallet(wallet *Wallet) error {
	_, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/wallet/open",
		makeJSONString(map[string]interface{}{
			"daemonHost": wallet.DaemonHost,
			"daemonPort": wallet.DaemonPort,
			"filename":   wallet.Filename,
			"password":   wallet.Password,
		}),
	)
	return err
}

// ImportKey - imports a wallet with a private spend and view key
func (wAPI WalletAPI) ImportKey(wallet *Wallet, viewKey, spendKey string, scanHeight uint64) error {
	_, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/wallet/import/key",
		makeJSONString(map[string]interface{}{
			"daemonHost":      wallet.DaemonHost,
			"daemonPort":      wallet.DaemonPort,
			"filename":        wallet.Filename,
			"password":        wallet.Password,
			"scanHeight":      scanHeight,
			"privateViewKey":  viewKey,
			"privateSpendKey": spendKey,
		}),
	)
	return err
}

// ImportSeed - imports a wallet using a mnemonic seed
func (wAPI WalletAPI) ImportSeed(wallet *Wallet, mnemonicSeed string, scanHeight uint64) error {
	_, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/wallet/import/seed",
		makeJSONString(map[string]interface{}{
			"daemonHost":   wallet.DaemonHost,
			"daemonPort":   wallet.DaemonPort,
			"filename":     wallet.Filename,
			"password":     wallet.Password,
			"scanHeight":   scanHeight,
			"mnemonicSeed": mnemonicSeed,
		}),
	)
	return err
}

// ImportView - imports a view only wallet using
func (wAPI WalletAPI) ImportView(wallet *Wallet, viewKey, address string, scanHeight uint64) error {
	_, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/wallet/import/view",
		makeJSONString(map[string]interface{}{
			"daemonHost":     wallet.DaemonHost,
			"daemonPort":     wallet.DaemonPort,
			"filename":       wallet.Filename,
			"password":       wallet.Password,
			"scanHeight":     scanHeight,
			"privateViewKey": viewKey,
			"address":        address,
		}),
	)
	return err
}

// CreateWallet - creates a new wallet
func (wAPI WalletAPI) CreateWallet(wallet *Wallet) error {
	_, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/wallet/create",
		makeJSONString(map[string]interface{}{
			"daemonHost": wallet.DaemonHost,
			"daemonPort": wallet.DaemonPort,
			"filename":   wallet.Filename,
			"password":   wallet.Password,
		}),
	)
	return err
}

// CloseWallet - saves and closes the wallet
func (wAPI WalletAPI) CloseWallet() error {
	_, _, err := wAPI.sendRequest(
		"DELETE",
		wAPI.Host+":"+wAPI.Port+"/wallet",
		"",
	)
	return err
}
