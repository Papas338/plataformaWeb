package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"

	"go.mongodb.org/mongo-driver/bson"
)

/*ModificarImagenes permite cambiar el registro en la bd */
func ModificarImagenes(ID string, tipo string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("imagenes")

	imagenes := make(map[string]interface{})

	var imagen models.Imagenes

	if tipo == "usuarioId" {
		condicion := bson.M{
			"usuarioId": "prueba",
		}

		err := col.FindOne(ctx, condicion).Decode(&imagen)
		if err != nil {
			fmt.Println("Registro no encontrado " + err.Error())
		}

		imagenes["usuarioId"] = ID

		updtString := bson.M{
			"$set": imagenes,
		}

		filtro := bson.M{"usuarioId": bson.M{"$eq": "prueba"}}

		_, err = col.UpdateOne(ctx, filtro, updtString)
		if err != nil {
			return false, err
		}

		return true, nil
	} else {
		condicion := bson.M{
			"usuarioId": ID,
		}

		err := col.FindOne(ctx, condicion).Decode(&imagen)
		if err != nil {
			fmt.Println("Registro no encontrado " + err.Error())
		}

		imagenes["imagen"+tipo] = ID + tipo + ".jpg"

		updtString := bson.M{
			"$set": imagenes,
		}

		filtro := bson.M{"usuarioId": bson.M{"$eq": ID}}

		_, err = col.UpdateOne(ctx, filtro, updtString)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}
