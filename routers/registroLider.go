package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
	"github.com/proyLSIPAZUD/plataformaWeb/blockchain"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*RegistroLider es la función para crear en la bd el registro de lider */
func RegistroLider(w http.ResponseWriter, r *http.Request) {
	var t models.LiderGeneral
	decodificar := json.NewDecoder(r.Body)
	err := decodificar.Decode(&t)

	//RegistroEsri(t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Nombre) == 0 {
		http.Error(w, "El nombre del lider es requerido", 400)
		return
	}

	//Se firma la transaccion que genero el pasante
	_, resultado := bd.ObtenerAdmin()

	key := []byte("example key 1234")
	adminPrivateKey := bd.Desencriptar(key, resultado.PrivateKey)

	hashSigner, signerAddress, _ := blockchain.SignerTransaction(adminPrivateKey, t.Hash)

	//Se regista la transaccion de la firma
	var firma models.Transaction

	firma.Tipo = "SignerTransaction"
	firma.Signer = signerAddress
	firma.Hash = hashSigner

	_, saveTransaction, _ := bd.RegistrarTransaccion(firma)
	if saveTransaction == false {
		http.Error(w, "Error al generar la transacción", 400)
		return
	}

	//Se genera la cuenta de blockchain para el lider social
	address, privateKey, publicKey := blockchain.GenerateRandom64HexString()
	_, statusPublicKey, _ := bd.VerificarPublicKey(publicKey)

	for statusPublicKey {
		address, privateKey, publicKey = blockchain.GenerateRandom64HexString()
		_, statusPublicKey, _ = bd.VerificarPublicKey(publicKey)
	}

	t.Address = address
	t.PrivateKey = privateKey
	t.PublicKey = publicKey

	fmt.Println("lider " + t.PrivateKey)

	//Se registra la transaccion de crear la cuenta del lider
	var create models.Transaction

	create.Tipo = "CreateLiderAccount"
	create.Signer = address

	_, saveTransaction, _ = bd.RegistrarTransaccion(create)
	if saveTransaction == false {
		http.Error(w, "Error al generar la transacción", 400)
		return
	}

	//Se hace una transferencia a la cuenta del lider
	_, resultado = bd.ObtenerAdmin()
	key = []byte("example key 1234")

	sender := bd.Desencriptar(key, resultado.PrivateKey)
	addresee := t.PrivateKey

	hashTransfer, cosigner, signer := blockchain.Transaction(sender, addresee)

	//Se registra la transaccion de la transferencia
	var transfer models.Transaction

	transfer.Tipo = "TransferLider"
	transfer.Hash = hashTransfer
	transfer.Cosigner = cosigner
	transfer.Signer = signer

	_, saveTransaction, _ = bd.RegistrarTransaccion(transfer)
	if saveTransaction == false {
		http.Error(w, "Error al generar la transacción", 400)
		return
	}

	//Se crea un namespace con la cuenta del lider
	HashNamespace := blockchain.RegisterNamespace(t.PrivateKey, strings.ReplaceAll(t.Nombre, " ", "_"))

	//Se registra la transaccion del namespace
	var namespace models.Transaction

	namespace.Tipo = "NamespaceLider"
	namespace.Hash = HashNamespace
	namespace.Signer = t.Address

	_, saveTransaction, _ = bd.RegistrarTransaccion(namespace)
	if saveTransaction == false {
		http.Error(w, "Error al generar la transacción", 400)
		return
	}

	//Se crea un alias con el namespace
	HashAlias := blockchain.TransactionAlias(t.PrivateKey, strings.ReplaceAll(t.Nombre, " ", "_"))

	fmt.Println("private key lider " + t.PrivateKey)

	//Se registra la transaccion del alias
	var alias models.Transaction

	alias.Tipo = "AliasLider"
	alias.Hash = HashAlias
	alias.Signer = t.Address

	_, saveTransaction, _ = bd.RegistrarTransaccion(alias)
	if saveTransaction == false {
		http.Error(w, "Error al generar la transacción", 400)
		return
	}

	var status bool
	var idUsuario string

	idUsuario, status, err = bd.InsertoLiderVerificado(t)
	t.UsuarioID = idUsuario[10 : len(idUsuario)-2]
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del lider resumen "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del lider resumen", 400)
		return
	}

	/* _, status, err = bd.InsertoLabor(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del lider labor "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del lider labor ", 400)
		return
	}

	_, status, err = bd.InsertoPrensa(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del lider prensa "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del lider prensa ", 400)
		return
	}

	_, status, err = bd.InsertoMedidasReparacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del lider medidas de reparacion "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del lider medidas reparacion ", 400)
		return
	}

	_, status, err = bd.InsertoIntereses(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del lider intereses "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del lider intereses ", 400)
		return
	}

	_, status, err = bd.InsertoReaccionTerritorio(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del lider reaccion territorio "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del lider reaccion territorio ", 400)
		return
	}

	_, status, err = bd.InsertoPalabrasClaves(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del lider palabras claves "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del lider palabras claves ", 400)
		return
	} */

	w.WriteHeader(http.StatusCreated)
}
