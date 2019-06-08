package walletapi

import "encoding/json"

// Status - represents a status object
type Status struct {
	NetworkBlockCount uint64 `json:"networkBlockCount"`
	WalletBlockCount  uint64 `json:"walletBlockCount"`
	LocalBlockCount   uint64 `json:"localDaemonBlockCount"`
	PeerCount         uint64 `json:"peerCount"`
	Hashrate          uint64 `json:"hashrate"`
	IsViewWallet      bool   `json:"isViewWallet"`
	SubWalletCount    uint64 `json:"subWalletCount"`
}

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

// ValidateAddress - validates an address
func (wAPI WalletAPI) ValidateAddress(address string) (*map[string]interface{}, error) {
	resp, _, err := wAPI.sendRequest(
		"POST",
		wAPI.Host+":"+wAPI.Port+"/addresses/validate",
		makeJSONString(map[string]interface{}{
			"address": address,
		}),
	)

	return resp, err
}

// Status - gets the wallet status
func (wAPI WalletAPI) Status() (*Status, error) {
	var stat Status
	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/status",
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &stat)
	}

	return &stat, err
}
