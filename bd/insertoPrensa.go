package bd

import (
	"context"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoPrensa es la parada final con la BD para insertar los datos del usuario */
func InsertoPrensa(t models.LiderGeneral) (string, bool, error) {
	var u models.Prensa

	u.UsuarioID = t.UsuarioID
	u.DetallePrensa = t.DetallePrensa
	u.DescripcionPrensa1 = t.DescripcionPrensa1
	u.DescripcionPrensa2 = t.DescripcionPrensa2
	u.DescripcionPrensa3 = t.DescripcionPrensa3
	u.DescripcionPrensa4 = t.DescripcionPrensa4
	u.DescripcionPrensa5 = t.DescripcionPrensa5
	u.LinkPrensa1 = t.LinkPrensa1
	u.LinkPrensa2 = t.LinkPrensa2
	u.LinkPrensa3 = t.LinkPrensa3
	u.LinkPrensa4 = t.LinkPrensa4
	u.LinkPrensa5 = t.LinkPrensa5

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("liderPrensa")

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
