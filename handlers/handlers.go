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

	router.HandleFunc("/insertoLider", middlew.ChequeoBD(middlew.ValidoJWT(routers.InsertoLider))).Methods("POST")
	router.HandleFunc("/registroLider", middlew.ChequeoBD(middlew.ValidoJWT(routers.RegistroLider))).Methods("POST")

	router.HandleFunc("/subirImagen", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirImagen))).Methods("POST")
	router.HandleFunc("/modificarImagenes", middlew.ChequeoBD(middlew.ValidoJWT(routers.ManejoImagenes))).Methods("PUT")
	router.HandleFunc("/obtenerImagenes", middlew.ChequeoBD(middlew.ValidoJWT(routers.ObtenerImagenes))).Methods("GET")

	router.HandleFunc("/listaLideresVerificar", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaLideresTemp))).Methods("GET")
	router.HandleFunc("/borrarLiderTemp", middlew.ChequeoBD(middlew.ValidoJWT(routers.BorrarLiderTemp))).Methods("DELETE")
	router.HandleFunc("/listaLideres", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaLideres))).Methods("GET")

	router.HandleFunc("/departamentos", middlew.ChequeoBD(middlew.ValidoJWT(routers.Departamentos))).Methods("GET")
	router.HandleFunc("/municipios", middlew.ChequeoBD(middlew.ValidoJWT(routers.Municipios))).Methods("GET")

	router.HandleFunc("/descargaInfo", middlew.ChequeoBD(middlew.ValidoJWT(routers.DescargaInfo))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
