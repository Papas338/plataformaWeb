package routers

import (
	"encoding/json"
	"net/http"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*ObtenerImagenes lee la lista de los imagenes */
func ObtenerImagenes(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var result models.Imagenes
	var status bool

	result, status, _ = bd.ConsultoImagenes(ID)
	if status == false {
		http.Error(w, "Error al leer registro de imagenes", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
