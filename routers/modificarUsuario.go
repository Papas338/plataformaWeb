package routers

import (
	"fmt"
	"net/http"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
	"github.com/proyLSIPAZUD/plataformaWeb/blockchain"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*ModificarUsuario modifica el role del usuario */
func ModificarUsuario(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	var users models.Usuario

	_, admPrivateKey, multiSignPublicKey := bd.ObtenerAdmin()

	key := []byte("example key 1234")
	adminPrivateKey := bd.Desencriptar(key, admPrivateKey)

	fmt.Println("Cuenta multifirma " + multiSignPublicKey)

	users, _ = bd.ObtenerUsuario(ID)

	userPrivateKey := bd.Desencriptar(key, users.PrivateKey)

	fmt.Println(users.Address)

	fmt.Println("Antes de transacción")
	fmt.Println(users.Role)

	if users.Role == "Usuario" {
		blockchain.AddCosigner(multiSignPublicKey, userPrivateKey, adminPrivateKey)
	} else {
		blockchain.RemoveCosigner(multiSignPublicKey, userPrivateKey, adminPrivateKey)
	}

	fmt.Println("Despues de transacción")

	status, err := bd.ModificoUsuario(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el role"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se actualizo el role", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
