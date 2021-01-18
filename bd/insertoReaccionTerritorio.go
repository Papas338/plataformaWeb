package bd

import (
	"context"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoReaccionTerritorio es la parada final con la BD para insertar los datos del usuario */
func InsertoReaccionTerritorio(t models.LiderGeneral) (string, bool, error) {
	var u models.ReaccionTerritorio

	u.UsuarioID = t.UsuarioID
	u.DetalleReaccionTerritorio = t.DetalleReaccionTerritorio
	u.DescripcionTerritorio = t.DescripcionTerritorio
	u.LinkTerritorio1 = t.LinkTerritorio1
	u.LinkTerritorio2 = t.LinkTerritorio2
	u.LinkTerritorio3 = t.LinkTerritorio3
	u.LinkTerritorio4 = t.LinkTerritorio4
	u.LinkTerritorio5 = t.LinkTerritorio5

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("liderReaccionTerritorio")

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
