package routers

import (
	"encoding/json"
	"net/http"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
	"github.com/proyLSIPAZUD/plataformaWeb/blockchain"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*Registro es la función para crear en la bd el registro de usuario */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
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

	_, adminPrivateKey, _ := bd.ObtenerAdmin()
	key := []byte("example key 1234")

	blockchain.Transaction(bd.Desencriptar(key, adminPrivateKey), t.PrivateKey)

	if t.Role != "Administrador" && t.Role != "Usuario" {
		http.Error(w, "Debe ingresar una rol valido", 400)
		return
	}

	if t.IsValidate != "Yes" && t.IsValidate != "No" {
		http.Error(w, "Debe ingresar una opción valida", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "ya existe un usuario registrado con ese email", 400)
	}

	_, status, err := bd.InsertoRegistro(t)
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
