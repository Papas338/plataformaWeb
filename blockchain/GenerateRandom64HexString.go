package blockchain

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

/*GenerateRandom64HexString se encarga de generar número hexadecimal para la creación de cuenta*/
func GenerateRandom64HexString() (string, string, string) {

	val, _ := randomHex(32)
	//fmt.Println(val)
	//Create account from private key
	aliceAccount, _ := sdk.NewAccountFromPrivateKey(val, sdk.PublicTest, &sdk.Hash{})
	PublicKeyCod := aliceAccount.KeyPair.PublicKey.Raw
	PublicKeyDec := hex.EncodeToString(PublicKeyCod)

	/* fmt.Printf("Address:\t%v\n", aliceAccount.Address)
	fmt.Printf("PrivateKey:\t%x\n", aliceAccount.KeyPair.PrivateKey.Raw)
	fmt.Printf("PublicKey:\t%x", aliceAccount.KeyPair.PublicKey.Raw) */

	return aliceAccount.Address.Address, val, PublicKeyDec
}
