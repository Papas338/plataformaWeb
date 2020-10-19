package bd

import (
	"context"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoPalabrasClaves es la parada final con la BD para insertar los datos del usuario */
func InsertoPalabrasClaves(t models.LiderGeneral) (string, bool, error) {
	var u models.PalabrasClaves

	u.UsuarioID = t.UsuarioID
	u.Palabra1 = t.Palabra1
	u.Palabra2 = t.Palabra2
	u.Palabra3 = t.Palabra3
	u.Palabra4 = t.Palabra4
	u.Palabra5 = t.Palabra5
	u.Palabra6 = t.Palabra6
	u.Palabra7 = t.Palabra7
	u.Palabra8 = t.Palabra8
	u.Palabra9 = t.Palabra9
	u.Palabra10 = t.Palabra10

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("liderPalabrasClaves")

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
