package routers

import (
	"fmt"
	"net/http"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
	"github.com/proyLSIPAZUD/plataformaWeb/blockchain"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*DescargaInfo registra la transaccion de la descarga de la info de un lider */
func DescargaInfo(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	key := []byte("example key 1234")

	var users models.Usuario

	users, _ = bd.ObtenerUsuario(IDUsuario)

	sender := bd.Desencriptar(key, users.PrivateKey)

	var lideres models.Lider

	lideres, _ = bd.ObtenerLider(ID)

	addresee := bd.Desencriptar(key, lideres.PrivateKey)

	fmt.Println(addresee)

	hashTransfer, cosigner, signer := blockchain.Transaction(sender, addresee)

	var transfer models.Transaction

	transfer.Tipo = "Transfer"
	transfer.Hash = hashTransfer
	transfer.Cosigner = cosigner
	transfer.Signer = signer

	_, saveTransaction, _ := bd.RegistrarTransaccion(transfer)
	if saveTransaction == false {
		http.Error(w, "Error al generar la transacción", 400)
		return
	}

	fmt.Println("Descarga realizada")

	w.WriteHeader(http.StatusCreated)
}
