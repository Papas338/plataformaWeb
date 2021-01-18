package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ObtenerAdmin permite evitar claves privadas iguales dentro de la blockchain */
func ObtenerAdmin() (bool, models.Usuario) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("usuarios")

	condicion := bson.M{"role": "Administrador"}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false, resultado
	}

	fmt.Println(resultado.MsAddress)

	return true, resultado
}
