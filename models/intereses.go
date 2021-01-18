package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Intereses es el modelo de información en la base de datos de MongoDB */
type Intereses struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID            string             `bson:"usuarioId" json:"usuarioId"`
	DetalleIntereses     string             `bson:"detalleIntereses" json:"detalleIntereses,omitempty"`
	DescripcionIntereses string             `bson:"descripcionIntereses" json:"descripcionIntereses,omitempty"`
	LinkIntereses1       string             `bson:"linkIntereses1" json:"linkIntereses1,omitempty"`
	LinkIntereses2       string             `bson:"linkIntereses2" json:"linkIntereses2,omitempty"`
	LinkIntereses3       string             `bson:"linkIntereses3" json:"linkIntereses3,omitempty"`
	LinkIntereses4       string             `bson:"linkIntereses4" json:"linkIntereses4,omitempty"`
	LinkIntereses5       string             `bson:"linkIntereses5" json:"linkIntereses5,omitempty"`
}
