package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*BuscarMunicipio encuentra el Municipio ingresado por el usuario */
func BuscarMunicipio(departamento string, municipio string) ([]*models.Municipio, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("municipios")

	var results []*models.Municipio

	findOptions := options.Find()

	query := bson.M{
		"NOM_DPTO": bson.M{"$regex": `(?i)` + departamento},
		"NOM_MPIO": bson.M{"$regex": `(?i)` + municipio},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	for cur.Next(ctx) {
		var s models.Municipio
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		results = append(results, &s)

	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
