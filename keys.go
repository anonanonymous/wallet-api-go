package walletapi

// ViewKey - gets the private view key of the wallet container
func (wAPI WalletAPI) ViewKey() (string, error) {
	var vk string

	resp, _, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/keys",
		"",
	)

	if err == nil {
		vk = (*resp)["privateViewKey"].(string)
	}

	return vk, err
}

// GetKeys - gets the public and private view key for the given address
func (wAPI WalletAPI) GetKeys(address string) (map[string]string, error) {
	keys := make(map[string]string, 2)

	resp, _, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/keys/"+address,
		"",
	)

	if err == nil {
		keys["privateSpendKey"] = (*resp)["privateSpendKey"].(string)
		keys["publicSpendKey"] = (*resp)["publicSpendKey"].(string)
	}

	return keys, err
}

// GetMnemonic - gets the mnemonic seed for the given address
func (wAPI WalletAPI) GetMnemonic(address string) (string, error) {
	var seed string

	resp, _, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/keys/mnemonic/"+address,
		"",
	)

	if err == nil {
		seed = (*resp)["mnemonicSeed"].(string)
	}

	return seed, err
}
