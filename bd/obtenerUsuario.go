package bd

import (
	"context"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ObtenerUsuario trae un usuario registrado en el sistema */
func ObtenerUsuario(ID string) (models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("usuarios")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return resultado, false
	}

	return resultado, true
}
