package walletapi

import (
	"fmt"
	"os"
	"testing"
)

var (
	W  *WalletAPI
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

	fmt.Println("====CreateWallet====")
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
	fmt.Println("Created address:", addr)
}

func TestStatus(t *testing.T) {
	stat, err := W.Status()
	if err != nil {
		t.Fail()
	}
	fmt.Println(stat)
}

func TestGetAddressBalance(t *testing.T) {
	balance, err := W.GetAddressBalance(addr)
	if err != nil {
		t.Fail()
	}
	fmt.Println(balance)
}

func TestClose(t *testing.T) {
	if err := W.CloseWallet(); err != nil {
		t.Fail()
	}
	os.Remove(wd + "/test.wallet")
	fmt.Println("====Done====")
}
