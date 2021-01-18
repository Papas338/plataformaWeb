package bd

import (
	"context"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ObtenerLider trae un lider registrado en el sistema */
func ObtenerLider(ID string) (models.Lider, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("lider")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	var resultado models.Lider

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return resultado, false
	}

	return resultado, true
}
