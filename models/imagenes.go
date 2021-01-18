package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Imagenes es el modelo de información en la base de datos de MongoDB */
type Imagenes struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID string             `bson:"usuarioId" json:"usuarioId"`
	Imagen1   string             `bson:"imagen1" json:"imagen1,omitempty"`
	Imagen2   string             `bson:"imagen2" json:"imagen2,omitempty"`
	Imagen3   string             `bson:"imagen3" json:"imagen3,omitempty"`
	Imagen4   string             `bson:"imagen4" json:"imagen4,omitempty"`
	Imagen5   string             `bson:"imagen5" json:"imagen5,omitempty"`
	Imagen6   string             `bson:"imagen6" json:"imagen6,omitempty"`
	Imagen7   string             `bson:"imagen7" json:"imagen7,omitempty"`
	Imagen8   string             `bson:"imagen8" json:"imagen8,omitempty"`
	Imagen9   string             `bson:"imagen9" json:"imagen9,omitempty"`
	Imagen10  string             `bson:"imagen10" json:"imagen10,omitempty"`
}
