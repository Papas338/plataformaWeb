package routers

import (
	"encoding/json"
	"net/http"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*ModificarUsuario modifica el role del usuario */
func ModificarUsuario(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos" + err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoUsuario(t, IDUsuario)
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
