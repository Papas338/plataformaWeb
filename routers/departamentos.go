package routers

import (
	"encoding/json"
	"net/http"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*Departamentos lee la lista de los departamentos */
func Departamentos(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	var result []*models.Departamento
	var status bool

	result, status = bd.BuscarDepartamento(search)
	if status == false {
		http.Error(w, "Error al leer los departamentos", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
