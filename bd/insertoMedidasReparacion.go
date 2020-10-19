package bd

import (
	"context"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoMedidasReparacion es la parada final con la BD para insertar los datos del usuario */
func InsertoMedidasReparacion(t models.LiderGeneral) (string, bool, error) {
	var u models.MedidasReparacion

	u.UsuarioID = t.UsuarioID
	u.DetalleReparacion = t.DetalleReparacion
	u.DescripcionMedidas1 = t.DescripcionMedidas1
	u.DescripcionMedidas2 = t.DescripcionMedidas2
	u.DescripcionMedidas3 = t.DescripcionMedidas3
	u.DescripcionMedidas4 = t.DescripcionMedidas4
	u.DescripcionMedidas5 = t.DescripcionMedidas5
	u.LinkMedidas1 = t.LinkMedidas1
	u.LinkMedidas2 = t.LinkMedidas2
	u.LinkMedidas3 = t.LinkMedidas3
	u.LinkMedidas4 = t.LinkMedidas4
	u.LinkMedidas5 = t.LinkMedidas5

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("liderMedidasReparacion")

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
