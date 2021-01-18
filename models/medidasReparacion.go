package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*MedidasReparacion es el modelo de información en la base de datos de MongoDB */
type MedidasReparacion struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID           string             `bson:"usuarioId" json:"usuarioId"`
	DetalleReparacion   string             `bson:"detalleReparacion" json:"detalleReparacion,omitempty"`
	DescripcionMedidas1 string             `bson:"descripcionMedidas1" json:"descripcionMedidas1,omitempty"`
	LinkMedidas1        string             `bson:"linkMedidas1" json:"linkMedidas1,omitempty"`
	DescripcionMedidas2 string             `bson:"descripcionMedidas2" json:"descripcionMedidas2,omitempty"`
	LinkMedidas2        string             `bson:"linkMedidas2" json:"linkMedidas2,omitempty"`
	DescripcionMedidas3 string             `bson:"descripcionMedidas3" json:"descripcionMedidas3,omitempty"`
	LinkMedidas3        string             `bson:"linkMedidas3" json:"linkMedidas3,omitempty"`
	DescripcionMedidas4 string             `bson:"descripcionMedidas4" json:"descripcionMedidas4,omitempty"`
	LinkMedidas4        string             `bson:"linkMedidas4" json:"linkMedidas4,omitempty"`
	DescripcionMedidas5 string             `bson:"descripcionMedidas5" json:"descripcionMedidas5,omitempty"`
	LinkMedidas5        string             `bson:"linkMedidas5" json:"linkMedidas5,omitempty"`
}
