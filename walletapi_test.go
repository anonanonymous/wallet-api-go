package walletapi

import (
	"os"
	"testing"
)

var (
	W        *WalletAPI
	addr, wd string
)

func TestSetup(t *testing.T) {
	wd, _ = os.Getwd()
	wallet := Wallet{
		Filename:   wd + "/test.wallet",
		Password:   "password",
		DaemonHost: "public.turtlenode.io",
		DaemonPort: 11898,
	}
	W = InitWalletAPI("password", "127.0.0.1", "8070")

	t.Log("====CreateWallet====")
	err := W.CreateWallet(&wallet)
	if err != nil {
		t.Fail()
	}
}

func TestCreateAddress(t *testing.T) {
	a, err := W.CreateAddress()
	if err != nil {
		t.Fail()
	}
	addr = a["address"]
	t.Log("Created address:", addr)
}

func TestNode(t *testing.T) {
	info, err := W.Node()
	if err != nil {
		t.Fail()
	}
	t.Log(info)
}
func TestStatus(t *testing.T) {
	stat, err := W.Status()
	if err != nil {
		t.Fail()
	}
	t.Log(stat)
}

func TestGetAddressBalance(t *testing.T) {
	balance, err := W.GetAddressBalance(addr)
	if err != nil {
		t.Fail()
	}
	t.Log(balance)
}

func TestClose(t *testing.T) {
	if err := W.CloseWallet(); err != nil {
		t.Fail()
	}
	os.Remove(wd + "/test.wallet")
	t.Log("====Done====")
}
