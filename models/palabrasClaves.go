package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*PalabrasClaves es el modelo de información en la base de datos de MongoDB */
type PalabrasClaves struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID string             `bson:"usuarioId" json:"usuarioId"`
	Palabra1  string             `bson:"palabra1" json:"palabra1,omitempty"`
	Palabra2  string             `bson:"palabra2" json:"palabra2,omitempty"`
	Palabra3  string             `bson:"palabra3" json:"palabra3,omitempty"`
	Palabra4  string             `bson:"palabra4" json:"palabra4,omitempty"`
	Palabra5  string             `bson:"palabra5" json:"palabra5,omitempty"`
	Palabra6  string             `bson:"palabra6" json:"palabra6,omitempty"`
	Palabra7  string             `bson:"palabra7" json:"palabra7,omitempty"`
	Palabra8  string             `bson:"palabra8" json:"palabra8,omitempty"`
	Palabra9  string             `bson:"palabra9" json:"palabra9,omitempty"`
	Palabra10 string             `bson:"palabra10" json:"palabra10,omitempty"`
}
