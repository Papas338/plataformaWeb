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

	_, resultado := bd.ObtenerAdmin()

	multiSignPublicKey := resultado.MsPublicKey

	key := []byte("example key 1234")
	adminPrivateKey := bd.Desencriptar(key, resultado.PrivateKey)

	fmt.Println("Cuenta multifirma " + multiSignPublicKey)

	users, _ = bd.ObtenerUsuario(ID)

	userPrivateKey := bd.Desencriptar(key, users.PrivateKey)

	fmt.Println(users.Address)

	fmt.Println("Antes de transacción")

	if users.Role == "Usuario" {
		hashBounded, hashLockFound, cosigner, admin := blockchain.AddCosigner(multiSignPublicKey, userPrivateKey, adminPrivateKey)
		var bounded models.Transaction

		bounded.Tipo = "AddCosigner"
		bounded.Hash = hashBounded
		bounded.Cosigner = cosigner

		_, saveTransaction, _ := bd.RegistrarTransaccion(bounded)
		if saveTransaction == false {
			http.Error(w, "No se agrego consignatario", 400)
			return
		}

		var lockFounds models.Transaction

		lockFounds.Tipo = "LockFounds"
		lockFounds.Hash = hashLockFound
		lockFounds.Cosigner = cosigner
		lockFounds.Signer = admin

		_, saveTransaction, _ = bd.RegistrarTransaccion(lockFounds)
		if saveTransaction == false {
			http.Error(w, "No se agrego consignatario por falta de fondos", 400)
			return
		}

	} else {
		hashCompleted, cosigner, admin := blockchain.RemoveCosigner(multiSignPublicKey, userPrivateKey, adminPrivateKey)
		var completed models.Transaction

		completed.Tipo = "RemoveCosigner"
		completed.Hash = hashCompleted
		completed.Cosigner = cosigner
		completed.Signer = admin

		_, saveTransaction, _ := bd.RegistrarTransaccion(completed)
		if saveTransaction == false {
			http.Error(w, "No se borro consignatario", 400)
			return
		}

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
