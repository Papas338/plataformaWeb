package routers

import (
	"encoding/json"
	"net/http"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
	"github.com/proyLSIPAZUD/plataformaWeb/blockchain"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*InsertoLider es la función para crear en la bd el registro del lider para verificación */
func InsertoLider(w http.ResponseWriter, r *http.Request) {
	var t models.LiderGeneral

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	//Se crea la transacción de registro de lider en blockchain
	var users models.Usuario

	_, resultado := bd.ObtenerAdmin()

	key := []byte("example key 1234")
	//adminPrivateKey := bd.Desencriptar(key, resultado.PrivateKey)
	multiSignPrivateKey := bd.Desencriptar(key, resultado.MsPrivateKey)

	users, _ = bd.ObtenerUsuario(IDUsuario)

	userPrivateKey := bd.Desencriptar(key, users.PrivateKey)

	hashBounded, _, cosigner := blockchain.LiderUploadTransaction(multiSignPrivateKey, userPrivateKey)

	var upload models.Transaction

	upload.Tipo = "uploadLider"
	upload.Hash = hashBounded
	upload.Cosigner = cosigner

	t.Hash = hashBounded

	_, saveTransaction, _ := bd.RegistrarTransaccion(upload)
	if saveTransaction == false {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del lider", 400)
		return
	}

	//Se inserta la información del lider en la base de datos
	_, status, err := bd.InsertoLiderGeneral(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
