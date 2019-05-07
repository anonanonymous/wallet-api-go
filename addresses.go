package walletapi

// Addresses - gets all the addresses in the wallet container
func (wAPI WalletAPI) Addresses() ([]string, error) {
	var addresses []string

	resp, _, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/addresses",
		"",
	)

	if err == nil {
		for _, v := range (*resp)["addresses"].([]interface{}) {
			addresses = append(addresses, v.(string))
		}
	}

	return addresses, err
}

// DeleteAddress -  deletes a subwallet in the wallet container
func (wAPI WalletAPI) DeleteAddress(address string) error {
	_, _, err := wAPI.sendRequest(
		"DELETE",
		wAPI.Host+":"+wAPI.Port+"/addresses/"+address,
		"",
	)
	return err
}

// PrimaryAddress - gets the primary address in the wallet container
func (wAPI WalletAPI) PrimaryAddress() (string, error) {
	resp, _, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/addresses/primary",
		"",
	)

	return (*resp)["address"].(string), err
}

// CreateAddress - creates a new random address in the wallet container
func (wAPI WalletAPI) CreateAddress() (map[string]string, error) {
	address := make(map[string]string, 3)

	resp, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/addresses/create",
		"",
	)

	if err == nil {
		address["address"] = (*resp)["address"].(string)
		address["privateSpendKey"] = (*resp)["privateSpendKey"].(string)
		address["publicSpendKey"] = (*resp)["publicSpendKey"].(string)
	}

	return address, err
}

// ImportAddress - import a subwallet with the given private spend key
func (wAPI WalletAPI) ImportAddress(spendKey string, scanHeight uint64) (string, error) {
	var address string

	resp, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/addresses/import",
		makeJSONString(map[string]interface{}{
			"scanHeight":      scanHeight,
			"privateSpendKey": spendKey,
		}),
	)

	if err == nil {
		address = (*resp)["address"].(string)
	}

	return address, err
}

// ImportViewAddress - import a view only subwallet with the given public spend key
func (wAPI WalletAPI) ImportViewAddress(spendKey string, scanHeight uint64) (string, error) {
	var address string

	resp, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/addresses/import/view",
		makeJSONString(map[string]interface{}{
			"scanHeight":     scanHeight,
			"publicSpendKey": spendKey,
		}),
	)

	if err == nil {
		address = (*resp)["address"].(string)
	}

	return address, err
}

// CreateIntegratedAddress - creates an integrated address from and address and paymentID
func (wAPI WalletAPI) CreateIntegratedAddress(address, paymentID string) (string, error) {
	var integratedAddress string

	resp, _, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/addresses/"+address+"/"+paymentID,
		"",
	)

	if err == nil {
		integratedAddress = (*resp)["integratedAddress"].(string)
	}

	return integratedAddress, err
}
