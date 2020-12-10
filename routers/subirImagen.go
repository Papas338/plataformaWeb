package routers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*SubirImagen sube la imagen al servidor */
func SubirImagen(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("imagen")
	var extension = strings.Split(handler.Filename, ".")[1]

	/* mensaje, indicador, status, err := bd.InsertoImagenes("prueba", extension)
	if err != nil {
		http.Error(w, mensaje+"Error al grabar la imagen en la BD ! "+err.Error(), http.StatusBadRequest)
		return
	} else if status == false {
		http.Error(w, mensaje, http.StatusBadRequest)
		return
	} */

	var archivo string = "plataformaWeb/lideres-sociales/src/assests/lideres/temp." + extension

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

	fmt.Println("imagen subida")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
