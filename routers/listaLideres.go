package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*ListaLideres lee la lista de los lideres */
func ListaLideres(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	var result []*models.Lider
	var status bool

	result, status = bd.VisualizarLiderResumen(pag, search)
	if status == false {
		http.Error(w, "Error al leer los lideres", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
