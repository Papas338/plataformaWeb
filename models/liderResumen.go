package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*LiderResumen es el modelo del resumen de información en la base de datos de MongoDB */
type LiderResumen struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre        string             `bson:"nombre" json:"nombre"`
	Fecha         string             `bson:"fecha" json:"fecha,omitempty"`
	Departamento  string             `bson:"departamento" json:"departamento,omitempty"`
	Municipio     string             `bson:"municipio" json:"municipio,omitempty"`
	TipoLiderazgo string             `bson:"tipoLiderazgo" json:"tipoLiderazgo,omitempty"`
	Territorio    string             `bson:"territorio" json:"territorio,omitempty"`
}
