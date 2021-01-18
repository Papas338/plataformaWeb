package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Departamento es el modelo de información en la base de datos de MongoDB */
type Departamento struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodDepartamento int                `bson:"COD_DPTO" json:"COD_DPTO"`
	NomDepartamento string             `bson:"NOM_DPTO" json:"NOM_DPTO,omitempty"`
	X               string             `bson:"X" json:"X,omitempty"`
	Y               string             `bson:"Y" json:"Y,omitempty"`
}
