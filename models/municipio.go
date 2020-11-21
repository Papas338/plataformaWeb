package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Municipio es el modelo de información en la base de datos de MongoDB */
type Municipio struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodDepartamento int                `bson:"COD_DPTO" json:"COD_DPTO"`
	NomDepartamento string             `bson:"NOM_DPTO" json:"NOM_DPTO,omitempty"`
	CodMunicipio    int                `bson:"COD_MPIO" json:"COD_MPIO"`
	NomMunicipio    string             `bson:"NOM_MPIO" json:"NOM_MPIO,omitempty"`
	Tipo            string             `bson:"TIPO" json:"TIPO,omiempty"`
	X               string             `bson:"X" json:"X,omitempty"`
	Y               string             `bson:"Y" json:"Y,omitempty"`
}
