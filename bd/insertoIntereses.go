package bd

import (
	"context"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoIntereses es la parada final con la BD para insertar los datos del usuario */
func InsertoIntereses(t models.LiderGeneral) (string, bool, error) {
	var u models.Intereses

	u.UsuarioID = t.UsuarioID
	u.DetalleIntereses = t.DetalleIntereses
	u.DescripcionIntereses = t.DescripcionIntereses
	u.LinkIntereses1 = t.LinkIntereses1
	u.LinkIntereses2 = t.LinkIntereses2
	u.LinkIntereses3 = t.LinkIntereses3
	u.LinkIntereses4 = t.LinkIntereses4
	u.LinkIntereses5 = t.LinkIntereses5

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("liderIntereses")

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
