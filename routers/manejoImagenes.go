package routers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
)

/*ManejoImagenes modifica el registro de las imagenes y mueve los archivos */
func ManejoImagenes(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	status, err := bd.ModificarImagenes(ID, "usuarioId")
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se actualizo el registro", 400)
		return
	}

	for i := 1; i <= 10; i++ {
		j := strconv.Itoa(i)
		OriginalPath := "uploads/lideres/prueba" + j + ".jpg"
		NewPath := "uploads/lideres/" + ID + j + ".jpg"
		e := os.Rename(OriginalPath, NewPath)
		if e == nil {
			bd.ModificarImagenes(ID, j)
		}
	}

	w.WriteHeader(http.StatusCreated)
}
