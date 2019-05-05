package walletapi

// Save - saves wallet container
func (wAPI WalletAPI) Save() error {
	_, _, err := wAPI.sendRequest(
		"PUT",
		wAPI.Host+":"+wAPI.Port+"/save",
		"",
	)

	return err
}

// Reset - resets and saves the wallet
func (wAPI WalletAPI) Reset(scanHeight uint64) error {
	_, _, err := wAPI.sendRequest(
		"PUT",
		wAPI.Host+":"+wAPI.Port+"/reset",
		makeJSONString(map[string]interface{}{
			"scanHeight": scanHeight,
		}),
	)

	return err
}

// Status - gets the wallet status
func (wAPI WalletAPI) Status() (*map[string]interface{}, error) {
	resp, _, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/status",
		"",
	)

	return resp, err
}
