package walletapi

import (
	"fmt"
	"os"
	"testing"
)

func TestSetup(t *testing.T) {
	wd, _ := os.Getwd()
	wallet := Wallet{
		Filename:   wd + "/test.wallet",
		Password:   "password",
		DaemonHost: "public.turtlenode.io",
		DaemonPort: 11898,
	}
	W := InitWalletAPI("password", "127.0.0.1", "8070")

	fmt.Println("====CreateWallet====")
	err := W.CreateWallet(&wallet)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if err := W.CloseWallet(); err != nil {
		t.Fail()
	}
	os.Remove(wd + "/test.wallet")
	fmt.Println("====Done====")
}
