package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario de la base de datos de MongoDB */
type Usuario struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"password,omitempty"`
	Role       string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	IsValidate string             `bson:"sitioWeb" json:"sitioWeb,omitempty"`
}
