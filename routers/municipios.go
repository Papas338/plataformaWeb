package routers

import (
	"encoding/json"
	"net/http"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*Municipios lee la lista de los Municipios */
func Municipios(w http.ResponseWriter, r *http.Request) {
	departamento := r.URL.Query().Get("departamento")
	municipio := r.URL.Query().Get("municipio")

	var result []*models.Municipio
	var status bool

	result, status = bd.BuscarMunicipio(departamento, municipio)
	if status == false {
		http.Error(w, "Error al leer los Municipios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
