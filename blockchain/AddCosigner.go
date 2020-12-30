package blockchain

import (
	"context"
	"fmt"
	"time"

	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
)

////Esta función es basada en AggregateBonded

/*AddCosigner permite agregar pasantes a las transacciones de blockchain */
func AddCosigner(multisigPublicKey string, cosignatoryPrivateKey string, adminPrivateKey string) {

	fmt.Println("addCosigner")

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
			{sdk.Add, cosigner.PublicAccount},
		},
	)
	if err != nil {
		fmt.Printf("NewModifyMultisigAccountTransaction returned error: %s", err)
		return
	}

	// Convert transactions to inner for an aggregate transaction
	transaction.ToAggregate(multisig)

	aggregateBondedTransaction, err := client.NewBondedAggregateTransaction(sdk.NewDeadline(time.Hour), []sdk.Transaction{transaction})

	// Sign transaction
	signedAggregateBoundedTransaction, err := cosigner.Sign(aggregateBondedTransaction)
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return
	}
	fmt.Printf("Content: \t\t%v", signedAggregateBoundedTransaction.Hash)

	// Create lock funds transaction for aggregate bounded
	lockFundsTransaction, err := client.NewLockFundsTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour),
		// Funds to lock
		sdk.XpxRelative(10),
		// Duration of lock transaction in blocks
		sdk.Duration(10),
		// Aggregate bounded transaction for lock
		signedAggregateBoundedTransaction,
	)
	if err != nil {
		fmt.Printf("NewLockFundsTransaction returned error: %s", err)
		return
	}

	// Sign transaction
	signedLockFundsTransaction, err := cosigner.Sign(lockFundsTransaction)
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return
	}
	fmt.Printf("Content: \t\t%v", signedLockFundsTransaction.Hash)

	// Announce transaction
	_, err = client.Transaction.Announce(context.Background(), signedLockFundsTransaction)
	if err != nil {
		fmt.Printf("Transaction.Announce returned error: %s", err)
		return
	}

	// Wait for lock funds transaction to be harvested
	time.Sleep(30 * time.Second)

	// Announce aggregate bounded transaction
	_, _ = client.Transaction.AnnounceAggregateBonded(context.Background(), signedAggregateBoundedTransaction)
	if err != nil {
		fmt.Printf("Transaction.AnnounceAggregateBonded returned error: %s", err)
		return
	}

	// Wait for aggregate bounded transaction to be harvested
	time.Sleep(30 * time.Second)

	/* // Create cosignature transaction from second account
	secondAccountCosignatureTransaction := sdk.NewCosignatureTransactionFromHash(signedAggregateBoundedTransaction.Hash)
	signedSecondAccountCosignatureTransaction, err := cosigner1.SignCosignatureTransaction(secondAccountCosignatureTransaction)
	if err != nil {
		fmt.Printf("SignCosignatureTransaction returned error: %s", err)
		return
	}

	// Announce transaction
	_, err = client.Transaction.AnnounceAggregateBondedCosignature(context.Background(), signedSecondAccountCosignatureTransaction)
	if err != nil {
		fmt.Printf("AnnounceAggregateBoundedCosignature returned error: %s", err)
		return
	} */

	// Create cosignature transaction from third account
	thirdAccountCosignatureTransaction := sdk.NewCosignatureTransactionFromHash(signedAggregateBoundedTransaction.Hash)
	signedThirdAccountCosignatureTransaction, err := admin.SignCosignatureTransaction(thirdAccountCosignatureTransaction)
	if err != nil {
		fmt.Printf("SignCosignatureTransaction returned error: %s", err)
		return
	}

	// Announce transaction
	_, err = client.Transaction.AnnounceAggregateBondedCosignature(context.Background(), signedThirdAccountCosignatureTransaction)
	if err != nil {
		fmt.Printf("AnnounceAggregateBoundedCosignature returned error: %s", err)
		return
	}
	// wait for the transaction to be confirmed! (very important)
	// you can use websockets to wait explicitly for transaction
	// to be in certain state, instead of hard waiting
	time.Sleep(45 * time.Second)
}
