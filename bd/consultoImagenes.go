package bd

import (
	"context"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoImagenes consulta la relacion entre 2 usuarios */
func ConsultoImagenes(ID string) (models.Imagenes, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("imagenes")

	condicion := bson.M{
		"usuarioId": ID,
	}

	var resultado models.Imagenes
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return resultado, false, err
	}
	return resultado, true, nil
}
