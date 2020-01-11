package walletapi

import (
	"os"
	"testing"
)

var (
	W             *WalletAPI
	addr, intAddr string
	wd            string
)

const r = "TRTLuySpDqd2fcvq5vx7Jiayw6yao7JHXFPuia5V83cVREtQSKyvWpxX9vamnUcG35BkQy6VfwUy5CsV9YNomioPGGyVhK3YXLq"

func TestSetup(t *testing.T) {
	//	wd, _ = os.Getwd()
	//	wallet := Wallet{
	//		Filename:   wd + "/test.wallet",
	//		Password:   "password",
	//		DaemonHost: "public.turtlenode.io",
	//		DaemonPort: 11898,
	//	}
	W = InitWalletAPI("pass", "127.0.0.1", "8070")

	//	err := W.CreateWallet(&wallet)
	//	if err != nil {
	//		t.Fail()
	//	}
	t.Log(W.Status())
	t.Log("Created Wallet")
}

func TestCreateAddress(t *testing.T) {
	a, err := W.CreateAddress()
	if err != nil {
		t.Error(err)
	}
	addr = a["address"]
	t.Log("Created address:", addr)
}

func TestCreateIntegratedAddress(t *testing.T) {
	tx := "c3fa5258221aeae7407ba3a2886811fb0c76ae1e8cdef179e1117d7ac6c9d3aa"
	a, err := W.CreateIntegratedAddress(addr, tx)
	if err != nil {
		t.Fail()
	}
	intAddr = a
	t.Log(a)
}

func TestValidateAddress(t *testing.T) {
	_, err := W.ValidateAddress(addr)
	if err != nil {
		t.Fail()
	}

	resp, err := W.ValidateAddress(intAddr)
	if err != nil {
		t.Fail()
	}

	if !(*resp).IsIntegrated {
		t.Fail()
	}
	t.Log(*resp)
}

/*
func TestPrepareTransactionBasic(t *testing.T) {
	tx, err := W.PrepareTransactionBasic(r, 100, "")
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	t.Log(W.DeletePreparedTransaction(tx.TransactionHash))
}
*
// run this test on a wallet with sufficient funds
func TestPrepareTransactionAdvanced(t *testing.T) {
	tx, err := W.PrepareTransactionAdvanced(
		[]map[string]interface{}{
			{
				"address": r,
				"amount":  10,
			},
		},
		nil, nil, nil, nil, "", "", nil,
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
	t.Log(W.DeletePreparedTransaction(tx.TransactionHash))

}


func TestSendTransaction(t *testing.T) {
	a := "TRTLv1N8cfS4wdGMZcqrct1J9eVznBbFMihQaUVNYSiSPRXuZddZny1VovkyLYZKzvGswmmL1j9gSiXcf7KWyHFke6DiyZztGkE"
	tx, err := W.SendTransactionAdvanced(
		[]map[string]interface{}{
			{
				"address": r,
				"amount":  10,
			},
		},
		nil, nil, nil, []string{a}, "", "", nil,
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(*tx)
}
*/

func TestGetKeys(t *testing.T) {
	pubKey, privKey, err := W.GetKeys(addr)
	if err != nil {
		t.Fail()
	}
	if pubKey == "" || privKey == "" {
		t.Fail()
	}
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

func TestGetBalance(t *testing.T) {
	unlocked, _, err := W.GetBalance()
	if err != nil {
		t.Error(err)
	}
	t.Log(unlocked)
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
