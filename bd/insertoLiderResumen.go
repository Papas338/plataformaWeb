package bd

import (
	"context"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoLiderResumen es la parada final con la BD para insertar los datos del usuario */
func InsertoLiderResumen(t models.LiderGeneral) (string, bool, error) {
	var u models.LiderResumen

	u.Nombre = t.Nombre
	u.Fecha = t.Fecha
	u.Departamento = t.Departamento
	u.Municipio = t.Municipio
	u.Territorio = t.Territorio
	u.TipoLiderazgo = t.TipoLiderazgo

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("liderResumen")

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
