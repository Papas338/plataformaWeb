package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/proyLSIPAZUD/plataformaWeb/middlew"
	"github.com/proyLSIPAZUD/plataformaWeb/routers"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el handler y pongo a escuchar el server */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/modificarUsuario", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarUsuario))).Methods("PUT")

	router.HandleFunc("/registroLider", middlew.ChequeoBD(middlew.ValidoJWT(routers.RegistroLider))).Methods("POST")

	router.HandleFunc("/subirImagen", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirImagen))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
