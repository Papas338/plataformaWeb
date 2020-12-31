package routers

import (
	"encoding/json"
	"net/http"

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

	var users models.Usuario

	_, resultado := bd.ObtenerAdmin()

	key := []byte("example key 1234")
	adminPrivateKey := bd.Desencriptar(key, resultado.PrivateKey)
	multiSignPrivateKey := bd.Desencriptar(key, resultado.MsPrivateKey)

	users, _ = bd.ObtenerUsuario("5fe810bd975220d567571ba7")

	userPrivateKey := bd.Desencriptar(key, users.PrivateKey)

	hashBounded, _, cosigner := blockchain.LiderUploadTransaction(multiSignPrivateKey, userPrivateKey, adminPrivateKey)

	var upload models.Transaction

	upload.Tipo = "uploadLider"
	upload.Hash = hashBounded
	upload.Cosigner = cosigner

	_, saveTransaction, _ := bd.RegistrarTransaccion(upload)
	if saveTransaction == false {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del lider", 400)
		return
	}

	address, privateKey, publicKey := blockchain.GenerateRandom64HexString()
	_, statusPublicKey, _ := bd.VerificarPublicKey(publicKey)

	for statusPublicKey {
		address, privateKey, publicKey = blockchain.GenerateRandom64HexString()
		_, statusPublicKey, _ = bd.VerificarPublicKey(publicKey)
	}

	t.Address = address
	t.PrivateKey = privateKey
	t.PublicKey = publicKey

	var create models.Transaction

	create.Tipo = "CreateLiderAccount"
	create.Signer = address

	_, resultado = bd.ObtenerAdmin()
	key = []byte("example key 1234")

	sender := bd.Desencriptar(key, resultado.PrivateKey)
	addresee := t.PrivateKey

	hashTransfer, cosigner, signer := blockchain.Transaction(sender, addresee)

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
