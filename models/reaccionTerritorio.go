package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ReaccionTerritorio es el modelo de información en la base de datos de MongoDB */
type ReaccionTerritorio struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID                 string             `bson:"usuarioId" json:"usuarioId"`
	DetalleReaccionTerritorio string             `bson:"detalleReaccionTerritorio" json:"detalleReaccionTerritorio,omitempty"`
	DescripcionTerritorio     string             `bson:"descripcionTerritorio" json:"descripcionTerritorio,omitempty"`
	LinkTerritorio1           string             `bson:"linkTerritorio1" json:"linkTerritorio1,omitempty"`
	LinkTerritorio2           string             `bson:"linkTerritorio2" json:"linkTerritorio2,omitempty"`
	LinkTerritorio3           string             `bson:"linkTerritorio3" json:"linkTerritorio3,omitempty"`
	LinkTerritorio4           string             `bson:"linkTerritorio4" json:"linkTerritorio4,omitempty"`
	LinkTerritorio5           string             `bson:"linkTerritorio5" json:"linkTerritorio5,omitempty"`
}
