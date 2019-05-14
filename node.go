package walletapi

import "encoding/json"

// NodeInfo - node response
type NodeInfo struct {
	DaemonHost  string `json:"daemonHost"`
	DaemonPort  int64  `json:"daemonPort"`
	NodeFee     int64  `json:"nodeFee"`
	NodeAddress string `json:"nodeAddress"`
}

// Node - gets the node address, port, fee, and fee address
func (wAPI WalletAPI) Node() (*NodeInfo, error) {
	var info NodeInfo
	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/node",
		"",
	)

	if err == nil {
		err = json.Unmarshal(*raw, &info)
	}

	return &info, err
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
