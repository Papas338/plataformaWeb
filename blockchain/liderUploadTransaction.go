package blockchain

import (
	"context"
	"fmt"
	"time"

	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
)

//Esta función es basada en AggregateBonded

/*LiderUploadTransaction permite crear transacciones de blockchain al cargar un lider social */
func LiderUploadTransaction(multisigPrivateKey string, cosignatoryPrivateKey string) (string, string, string) {

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return "", "", ""
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// Create an account from a private key
	multisig, err := client.NewAccountFromPrivateKey(multisigPrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPublicKey returned error: %s", err)
		return "", "", ""
	}

	cosigner, err := client.NewAccountFromPrivateKey(cosignatoryPrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return "", "", ""
	}

	// Create a new transfer type transaction
	transaction, err := client.NewTransferTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// The address of the recipient account.
		cosigner.Address,
		// The array of mosaic to be sent.
		[]*sdk.Mosaic{sdk.Xpx(10)},
		// The transaction message of 1024 characters.
		sdk.NewPlainMessage("Carga de información lider social"),
	)
	if err != nil {
		fmt.Printf("NewTransferTransaction returned error: %s", err)
		return "", "", ""
	}

	// Convert transactions to inner for an aggregate transaction
	transaction.ToAggregate(multisig.PublicAccount)

	aggregateBondedTransaction, err := client.NewBondedAggregateTransaction(sdk.NewDeadline(time.Hour), []sdk.Transaction{transaction})

	// Sign transaction
	signedAggregateBoundedTransaction, err := cosigner.Sign(aggregateBondedTransaction)
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return "", "", ""
	}
	hashBounded := fmt.Sprintf("%s", signedAggregateBoundedTransaction.Hash)

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
		return "", "", ""
	}

	// Sign transaction
	signedLockFundsTransaction, err := cosigner.Sign(lockFundsTransaction)
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return "", "", ""
	}
	hashLockFunds := fmt.Sprintf("%s", signedLockFundsTransaction.Hash)

	// Announce transaction
	_, err = client.Transaction.Announce(context.Background(), signedLockFundsTransaction)
	if err != nil {
		fmt.Printf("Transaction.Announce returned error: %s", err)
		return "", "", ""
	}

	// Wait for lock funds transaction to be harvested
	time.Sleep(30 * time.Second)

	// Announce aggregate bounded transaction
	_, _ = client.Transaction.AnnounceAggregateBonded(context.Background(), signedAggregateBoundedTransaction)
	if err != nil {
		fmt.Printf("Transaction.AnnounceAggregateBonded returned error: %s", err)
		return "", "", ""
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

	/* // Create cosignature transaction from third account
	thirdAccountCosignatureTransaction := sdk.NewCosignatureTransactionFromHash(signedAggregateBoundedTransaction.Hash)
	signedThirdAccountCosignatureTransaction, err := admin.SignCosignatureTransaction(thirdAccountCosignatureTransaction)
	if err != nil {
		fmt.Printf("SignCosignatureTransaction returned error: %s", err)
		return "", "", ""
	}

	// Announce transaction
	_, err = client.Transaction.AnnounceAggregateBondedCosignature(context.Background(), signedThirdAccountCosignatureTransaction)
	if err != nil {
		fmt.Printf("AnnounceAggregateBoundedCosignature returned error: %s", err)
		return "", "", ""
	}
	// wait for the transaction to be confirmed! (very important)
	// you can use websockets to wait explicitly for transaction
	// to be in certain state, instead of hard waiting
	time.Sleep(45 * time.Second) */

	return hashBounded, hashLockFunds, cosigner.Address.Address
}
