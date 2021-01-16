package blockchain

import (
	"context"
	"fmt"

	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
)

/*SignerTransaction permite firmar las transacciones por parte del administrador*/
func SignerTransaction(adminPrivateKey string, hash string) (string, string, string) {

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return "", "", ""
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	fmt.Println("Este es le hash" + hash)

	transactionStatus, err := client.Transaction.GetTransactionStatus(context.Background(), hash)
	if err != nil {
		fmt.Printf("Transaction.GetTransactionStatus returned error: %s", err)
		return "", "", ""
	}

	admin, err := client.NewAccountFromPrivateKey(adminPrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return "", "", ""
	}

	fmt.Println("Creando firma")
	// Create cosignature transaction from third account
	thirdAccountCosignatureTransaction := sdk.NewCosignatureTransactionFromHash(transactionStatus.Hash)
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
	fmt.Println("Terminando firma")

	return signedThirdAccountCosignatureTransaction.ParentHash.String(), admin.Address.Address, ""
}
