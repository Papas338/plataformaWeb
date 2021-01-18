package routers

import (
	"net/http"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
)

/*BorrarLiderTemp modifica el role del usuario */
func BorrarLiderTemp(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	status, err := bd.BorrarLideresTemp(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el lider"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se borro el lider", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
