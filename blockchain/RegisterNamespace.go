package blockchain

import (
	"context"
	"fmt"
	"time"

	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
)

/*RegisterNamespace crea un namespace para los lideres sociales */
func RegisterNamespace(privateKey string, namespace string) string {

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return ""
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// Create an account from a private key
	account, err := client.NewAccountFromPrivateKey(privateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return ""
	}

	// Create a new register root-Namespace type transaction
	transaction, err := client.NewRegisterRootNamespaceTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// The namespace Name.
		namespace,
		// The duration per block.
		sdk.Duration(10),
	)
	if err != nil {
		fmt.Printf("NewRegisterRootNamespaceTransaction returned error: %s", err)
		return ""
	}

	// Sign transaction
	signedTransaction, err := account.Sign(transaction)
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return ""
	}
	// Announce transaction
	_, err = client.Transaction.Announce(context.Background(), signedTransaction)
	if err != nil {
		fmt.Printf("Transaction.Announce returned error: %s", err)
		return ""
	}

	// wait for the transaction to be confirmed! (very important)
	// you can use websockets to wait explicitly for transaction
	// to be in certain state, instead of hard waiting
	time.Sleep(time.Second * 5)
	Hash := fmt.Sprintf("%s", signedTransaction.Hash)

	return Hash

}
