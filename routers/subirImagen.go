package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
)

/*SubirImagen sube la imagen al servidor */
func SubirImagen(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("imagen")
	var extension = strings.Split(handler.Filename, ".")[1]

	ID := "prueba"

	mensaje, indicador, status, err := bd.InsertoImagenes(ID, extension)
	if err != nil {
		http.Error(w, mensaje+"Error al grabar el avatar en la BD ! "+err.Error(), http.StatusBadRequest)
		return
	} else if status == false {
		http.Error(w, mensaje, http.StatusBadRequest)
		return
	}

	var archivo string = "uploads/lideres/" + ID + indicador + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
