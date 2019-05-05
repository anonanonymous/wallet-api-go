package walletapi

// WalletAPI - communicates with wallet-api
type WalletAPI struct {
	APIKey string
	Host   string
	Port   string
}

// InitWalletAPI - initializes wallet-api connection
func InitWalletAPI(apiKey, host, port string) *WalletAPI {
	return &WalletAPI{
		APIKey: apiKey,
		Host:   "http://" + host,
		Port:   port,
	}
}

// Constants
const (
	MIXIN = 3
	FEE   = 10
)
