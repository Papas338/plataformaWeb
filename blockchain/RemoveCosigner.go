package blockchain

import (
	"context"
	"fmt"
	"time"

	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
)

const (
	// Minimal approval count
	minimalApproval = 0
	// Minimal removal count
	minimalRemoval = 0
)

//Esta función es basada en AggregateComplete

/*RemoveCosigner permite eliminar al pasante de las transacciones en blockchain */
func RemoveCosigner(multisigPublicKey string, cosignatoryPrivateKey string, adminPrivateKey string) {

	fmt.Println("removeCosigner")

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// Create an account from a private key
	multisig, err := client.NewAccountFromPublicKey(multisigPublicKey)
	if err != nil {
		fmt.Printf("NewAccountFromPublicKey returned error: %s", err)
		return
	}

	cosigner, err := client.NewAccountFromPrivateKey(cosignatoryPrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}

	/* cosigner1, err := client.NewAccountFromPrivateKey(cosign1PrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	} */

	admin, err := client.NewAccountFromPrivateKey(adminPrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}

	transaction, err := client.NewModifyMultisigAccountTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// The number of signatures needed to approve a transaction.
		minimalApproval,
		// The number of signatures needed to remove a cosignatory.
		minimalRemoval,
		// Array of cosigner accounts added or removed from the multisignature account.
		[]*sdk.MultisigCosignatoryModification{
			{sdk.Remove, cosigner.PublicAccount},
		},
	)
	if err != nil {
		fmt.Printf("NewModifyMultisigAccountTransaction returned error: %s", err)
		return
	}

	// Convert transactions to inner for an aggregate transaction
	// De esta cuenta es que salen los fondos
	transaction.ToAggregate(multisig)

	aggregateCompletedTransaction, err := client.NewCompleteAggregateTransaction(sdk.NewDeadline(time.Hour), []sdk.Transaction{transaction})

	// Sign transaction
	signedAggregateCompletedTransaction, err := admin.Sign(aggregateCompletedTransaction)
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return
	}
	fmt.Printf("Content: \t\t%v", signedAggregateCompletedTransaction.Hash)

	// Announce aggregate bounded transaction
	_, _ = client.Transaction.Announce(context.Background(), signedAggregateCompletedTransaction)
	if err != nil {
		fmt.Printf("Transaction.AnnounceAggregateCompleted returned error: %s", err)
		return
	}

	// Wait for aggregate bounded transaction to be harvested
	time.Sleep(30 * time.Second)

}
