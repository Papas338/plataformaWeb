package bd

import (
	"context"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoLabor es la parada final con la BD para insertar los datos del usuario */
func InsertoLabor(t models.LiderGeneral) (string, bool, error) {
	var u models.RelacionLabor

	u.UsuarioID = t.UsuarioID
	u.DescripcionRelacionLabor = t.DescripcionRelacionLabor
	u.DetalleRelacionLabor = t.DetalleRelacionLabor
	u.LinkRelacionLabor1 = t.LinkRelacionLabor1
	u.LinkRelacionLabor2 = t.LinkRelacionLabor2
	u.LinkRelacionLabor3 = t.LinkRelacionLabor3
	u.LinkRelacionLabor4 = t.LinkRelacionLabor4
	u.LinkRelacionLabor5 = t.LinkRelacionLabor5

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("liderLabor")

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
