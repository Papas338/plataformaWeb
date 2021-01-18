package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Transaction es el modelo de transacción en blockchain de la base de datos de MongoDB */
type Transaction struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Tipo     string             `bson:"tipo" json:"tipo"`
	Hash     string             `bson:"hash" json:"hash,omitempty"`
	Cosigner string             `bson:"cosigner" json:"cosigner,omitempty"`
	Signer   string             `bson:"signer" json:"signer,omitempty"`
}
