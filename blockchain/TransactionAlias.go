package blockchain

import (
	"context"
	"fmt"
	"time"

	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
)

/*TransactionAlias dispone el alias al namespace del lider social */
func TransactionAlias(privateKey string, namespace string) string {

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

	// Create namespace id from it's name
	namespaceId, err := sdk.NewNamespaceIdFromName(namespace)
	if err != nil {
		fmt.Printf("NewNamespaceIdFromName returned error: %s", err)
		return ""
	}

	// Create a address alias transaction
	transaction, err := client.NewAddressAliasTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour),
		// The address which we want to link
		account.Address,
		// The namespace id which we want to link.
		namespaceId,
		// The type of action ( in this case we want to link entities ).
		sdk.AliasLink,
	)
	if err != nil {
		fmt.Printf("NewAddressAliasTransaction returned error: %s", err)
		return ""
	}

	// Sign transaction
	signedTransaction, err := account.Sign(transaction)
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return ""
	}
	Hash := fmt.Sprintf("%s", signedTransaction.Hash)

	// Announce transaction
	_, err = client.Transaction.Announce(context.Background(), signedTransaction)
	if err != nil {
		fmt.Printf("Transaction.Announce returned error: %s", err)
		return ""
	}

	return Hash
}
