package walletapi

// Node - gets the node address, port, fee, and fee address
func (wAPI WalletAPI) Node() (map[string]interface{}, error) {
	resp, _, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/node",
		"",
	)
	return *resp, err
}

// SetNode - sets the node address and port
func (wAPI WalletAPI) SetNode(daemonHost string, daemonPort int) error {
	_, _, err := wAPI.sendRequest(
		"PUT",
		wAPI.Host+":"+wAPI.Port+"/node",
		makeJSONString(map[string]interface{}{
			"daemonHost": daemonHost,
			"daemonPort": daemonPort,
		}),
	)
	return err
}
