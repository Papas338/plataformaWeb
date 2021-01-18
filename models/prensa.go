package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Prensa es el modelo de información en la base de datos de MongoDB */
type Prensa struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID          string             `bson:"usuarioId" json:"usuarioId"`
	DetallePrensa      string             `bson:"detallePrensa" json:"detallePrensa,omitempty"`
	DescripcionPrensa1 string             `bson:"descripcionPrensa1" json:"descripcionPrensa1,omitempty"`
	LinkPrensa1        string             `bson:"linkPrensa1" json:"linkPrensa1,omitempty"`
	DescripcionPrensa2 string             `bson:"descripcionPrensa2" json:"descripcionPrensa2,omitempty"`
	LinkPrensa2        string             `bson:"linkPrensa2" json:"linkPrensa2,omitempty"`
	DescripcionPrensa3 string             `bson:"descripcionPrensa3" json:"descripcionPrensa3,omitempty"`
	LinkPrensa3        string             `bson:"linkPrensa3" json:"linkPrensa3,omitempty"`
	DescripcionPrensa4 string             `bson:"descripcionPrensa4" json:"descripcionPrensa4,omitempty"`
	LinkPrensa4        string             `bson:"linkPrensa4" json:"linkPrensa4,omitempty"`
	DescripcionPrensa5 string             `bson:"descripcionPrensa5" json:"descripcionPrensa5,omitempty"`
	LinkPrensa5        string             `bson:"linkPrensa5" json:"linkPrensa5,omitempty"`
}
