package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*ModificoRegistro permite modificar el perfil del usuario */
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})
	if len(u.Role) > 0 {
		registro["role"] = u.Role
	}
	if len(u.IsValidate) > 0 {
		registro["isvalidate"] = u.IsValidate
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
