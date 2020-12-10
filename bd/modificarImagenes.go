package bd

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/*ModificarImagenes permite cambiar el registro en la bd */
func ModificarImagenes(ID string) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("proyectolideres")
	col := db.Collection("imagenes")

	imagenes, status, _ := ConsultoImagenes(ID)
	obtener := structToMap(&imagenes)
	registro := make(map[string]interface{})

	//imagenes := make(map[string]interface{})

	//var imagen models.Imagenes

	if status == false {
		fmt.Println("priemra imagen")
		/* condicion := bson.M{
			"usuarioId": "prueba",
		}

		err := col.FindOne(ctx, condicion).Decode(&imagen)
		if err != nil {
			fmt.Println("Registro no encontrado " + err.Error())
		}

		registro["usuarioId"] = ID

		updtString := bson.M{
			"$set": registro,
		}

		filtro := bson.M{"usuarioId": bson.M{"$eq": "prueba"}}

		_, err = col.UpdateOne(ctx, filtro, updtString)
		if err != nil {
			return "", false, err
		} */

		imagenes.UsuarioID = ID
		imagenes.Imagen1 = ID + "1.jpg"
		OriginalPath := "plataformaWeb/lideres-sociales/src/assests/lideres/temp.jpg"
		NewPath := "plataformaWeb/lideres-sociales/src/assests/lideres/" + ID + "1.jpg"
		e := os.Rename(OriginalPath, NewPath)
		if e != nil {
			fmt.Println("No encontro imagen 1")
		}
		_, err := col.InsertOne(ctx, imagenes)
		if err != nil {
			return "", false, err
		}

		return "", true, nil
	} else {
		fmt.Println("otra imagen")
		/* condicion := bson.M{
			"usuarioId": ID,
		}

		err := col.FindOne(ctx, condicion).Decode(&imagen)
		if err != nil {
			fmt.Println("Registro no encontrado " + err.Error())
		} */

		/* imagenes["imagen"+tipo] = ID + tipo + ".jpg"

		updtString := bson.M{
			"$set": imagenes,
		}

		filtro := bson.M{"usuarioId": bson.M{"$eq": ID}}

		_, err = col.UpdateOne(ctx, filtro, updtString)
		if err != nil {
			return false, err
		} */

		guardados := 0
		for i := 1; i <= 10; i++ {
			fmt.Println(obtener["Imagen"+strconv.Itoa(i)][0] + "prueba")
			if obtener["Imagen"+strconv.Itoa(i)][0] == "" {
				registro["imagen"+strconv.Itoa(i)] = ID + strconv.Itoa(i) + ".jpg"
				updtString := bson.M{
					"$set": registro,
				}

				filtro := bson.M{"usuarioId": bson.M{"$eq": ID}}

				_, err := col.UpdateOne(ctx, filtro, updtString)
				if err != nil {
					return "", false, err
				}

				j := strconv.Itoa(i)
				OriginalPath := "plataformaWeb/lideres-sociales/src/assests/lideres/temp.jpg"
				NewPath := "plataformaWeb/lideres-sociales/src/assests/lideres/" + ID + j + ".jpg"
				e := os.Rename(OriginalPath, NewPath)
				if e != nil {
					fmt.Println("No encontro imagen " + j)
				}

				return "Imagenes cargadas", true, nil
			} else {
				guardados++
			}
			if guardados == 10 {
				return "Limite de imagenes cumplido", false, nil
			}
		}

		return "", true, nil
	}
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
