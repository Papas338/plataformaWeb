package bd

import (
	"context"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoLiderVerificado es la parada final con la BD para insertar los datos del usuario */
func InsertoLiderVerificado(t models.LiderGeneral) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("lider")

	key := []byte("example key 1234")

	t.PrivateKey = Encriptar(key, t.PrivateKey)

	result, err := col.InsertOne(ctx, t)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
