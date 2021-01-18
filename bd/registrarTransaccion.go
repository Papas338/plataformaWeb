package bd

import (
	"context"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*RegistrarTransaccion es la parada final con la BD para insertar los datos de una transacción blockchain */
func RegistrarTransaccion(u models.Transaction) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("transacciones")

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
