# wallet-api-go
![Go Report Card](https://goreportcard.com/badge/github.com/anonanonymous/wallet-api-go)  
## Prerequisites
- Go compiler
- wallet-api daemon
## Getting Started
Start the wallet-api daemon.  
Then make and switch to directory named `test`  
```
mkdir test && cd $_
```
Create main.go with the following contents. Edit the wallet configuration if neccessary.  
```
package main

import (
	"fmt"
	"os"
	walletapi "github.com/anonanonymous/wallet-api-go"
)

func main() {
	wd, _ := os.Getwd()
	wallet := walletapi.Wallet{
		Filename:   wd+"/test.wallet",
		Password:   "password",
		DaemonHost: "public.turtlenode.io",
		DaemonPort: 11898,
	}
	W := walletapi.InitWalletAPI("password", "127.0.0.1", "8070")

	fmt.Println("====CreateWallet====")
	fmt.Println(W.CreateWallet(&wallet))

	fmt.Println("====Status====")
	fmt.Println(W.Status())

	fmt.Println("====Node====")
	fmt.Println(W.Node())

	fmt.Println("====ViewKey====")
	fmt.Println(W.ViewKey())

	fmt.Println("====PrimaryAddress===")
	primary, _ := W.PrimaryAddress()
	fmt.Println(primary)

	fmt.Println("====Addresses====")
	addresses, _ := W.Addresses()
	fmt.Println(addresses)

	fmt.Println("====ValidateAddress====")
	fmt.Println(W.ValidateAddress(primary))

	fmt.Println("====GetBalance====")
	fmt.Println(W.GetBalance())

	fmt.Println("====GetBalances====")
	fmt.Println(W.GetBalances())

	fmt.Println("====GetKeys====")
	fmt.Println(W.GetKeys(primary))

	fmt.Println("====GetMnemonic====")
	fmt.Println(W.GetMnemonic(primary))

	W.CloseWallet()
	fmt.Println("====Done====")
}
```
Now run `go mod init "github.com/you/test"` to initialize the module  
Build and run the example program 
`go build; ./test`  
