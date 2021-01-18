package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/proyLSIPAZUD/plataformaWeb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*VisualizarLiderResumen lee los lideres en el sistema */
func VisualizarLiderResumen(page int64, search string) ([]*models.Lider, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("lider")

	var results []*models.Lider

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 5)
	findOptions.SetLimit(5)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	for cur.Next(ctx) {
		var s models.Lider
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
