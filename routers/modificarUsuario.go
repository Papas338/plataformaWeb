package routers

import (
	"net/http"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
)

/*ModificarUsuario modifica el role del usuario */
func ModificarUsuario(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

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
