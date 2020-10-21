package bd

import (
	"context"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/*InsertoImagenes es la parada final con la BD para insertar los datos del usuario */
func InsertoImagenes(ID string, extension string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("imagenes")

	imagenes, status, _ := ConsultoImagenes(ID)
	obtener := structToMap(&imagenes)
	registro := make(map[string]interface{})

	if status == false {
		imagenes.UsuarioID = ID
		imagenes.Imagen1 = ID + "1." + extension
		_, err := col.InsertOne(ctx, imagenes)
		if err != nil {
			return "", false, err
		}
	} else {
		guardados := 0
		for i := 1; i <= 10; i++ {
			if obtener["Imagen"+strconv.Itoa(i)][0] == "" {
				registro["imagen"+strconv.Itoa(i)] = ID + strconv.Itoa(i) + "." + extension
				updtString := bson.M{
					"$set": registro,
				}

				filtro := bson.M{"usuarioId": bson.M{"$eq": ID}}

				_, err := col.UpdateOne(ctx, filtro, updtString)
				if err != nil {
					return "", false, err
				}

				return "ObjID.String()", true, nil
			} else {
				guardados++
			}
			if guardados == 10 {
				return "Limite de imagenes cumplido", false, nil
			}
		}
	}

	//ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return "Imagenes cargadas", true, nil
}

func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		// You ca use tags here...
		// tag := typ.Field(i).Tag.Get("tagname")
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(typ.Field(i).Name, v)
	}
	return
}
