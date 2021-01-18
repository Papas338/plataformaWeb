package routers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
)

/*ManejoImagenes modifica el registro de las imagenes y mueve los archivos */
func ManejoImagenes(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	time.Sleep(120 * time.Second)
	fmt.Println("sleep terminado")

	estado, status, err := bd.ModificarImagenes(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se actualizo el registro", 400)
		return
	}

	fmt.Println(estado)

	/* for i := 1; i <= 10; i++ {
		j := strconv.Itoa(i)
		OriginalPath := "uploads/lideres/prueba" + j + ".jpg"
		NewPath := "uploads/lideres/" + ID + j + ".jpg"
		e := os.Rename(OriginalPath, NewPath)
		if e == nil {
			bd.ModificarImagenes(ID, j)
		} else {
			fmt.Println("No encontro imagen " + j)
		}
	} */

	w.WriteHeader(http.StatusCreated)
}
