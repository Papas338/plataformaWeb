package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BorrarLideresTemp permite modificar el perfil del usuario */
func BorrarLideresTemp(ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("liderGeneral")

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.DeleteOne(ctx, filtro, nil)
	if err != nil {
		return false, err
	}

	return true, nil
}
