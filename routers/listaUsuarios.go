package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*ListaUsuarios lee la lista de los usuarios */
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	var result []*models.Usuario
	var status bool

	if RoleUsuario == "Administrador" {
		result, status = bd.VisualizarUsuarios(pag, search)
		if status == false {
			http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "No puede visualizar usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
