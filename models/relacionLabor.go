package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*RelacionLabor es el modelo de información en la base de datos de MongoDB */
type RelacionLabor struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID                string             `bson:"usuarioId" json:"usuarioId"`
	DetalleRelacionLabor     string             `bson:"detalleRelacionLabor" json:"detalleRelacionLabor,omitempty"`
	DescripcionRelacionLabor string             `bson:"descripcionRelacionLabor" json:"descripcionRelacionLabor,omitempty"`
	LinkRelacionLabor1       string             `bson:"linkRelacionLabor1" json:"linkRelacionLabor1,omitempty"`
	LinkRelacionLabor2       string             `bson:"linkRelacionLabor2" json:"linkRelacionLabor2,omitempty"`
	LinkRelacionLabor3       string             `bson:"linkRelacionLabor3" json:"linkRelacionLabor3,omitempty"`
	LinkRelacionLabor4       string             `bson:"linkRelacionLabor4" json:"linkRelacionLabor4,omitempty"`
	LinkRelacionLabor5       string             `bson:"linkRelacionLabor5" json:"linkRelacionLabor5,omitempty"`
}
