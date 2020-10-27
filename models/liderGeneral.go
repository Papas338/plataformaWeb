package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*LiderGeneral es el modelo de información en la base de datos de MongoDB */
type LiderGeneral struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre                    string             `bson:"nombre" json:"nombre"`
	Fecha                     string             `bson:"fecha" json:"fecha,omitempty"`
	Departamento              string             `bson:"departamento" json:"departamento,omitempty"`
	Municipio                 string             `bson:"municipio" json:"municipio,omitempty"`
	TipoLiderazgo             string             `bson:"tipoLiderazgo" json:"tipoLiderazgo,omitempty"`
	Territorio                string             `bson:"territorio" json:"territorio,omitempty"`
	UsuarioID                 string             `bson:"usuarioId" json:"usuarioId"`
	DetalleIntereses          string             `bson:"detalleIntereses" json:"detalleIntereses,omitempty"`
	DescripcionIntereses      string             `bson:"descripcionIntereses" json:"descripcionIntereses,omitempty"`
	LinkIntereses1            string             `bson:"linkIntereses1" json:"linkIntereses1,omitempty"`
	LinkIntereses2            string             `bson:"linkIntereses2" json:"linkIntereses2,omitempty"`
	LinkIntereses3            string             `bson:"linkIntereses3" json:"linkIntereses3,omitempty"`
	LinkIntereses4            string             `bson:"linkIntereses4" json:"linkIntereses4,omitempty"`
	LinkIntereses5            string             `bson:"linkIntereses5" json:"linkIntereses5,omitempty"`
	DetalleReparacion         string             `bson:"detalleReparacion" json:"detalleReparacion,omitempty"`
	DescripcionMedidas1       string             `bson:"descripcionMedidas1" json:"descripcionMedidas1,omitempty"`
	LinkMedidas1              string             `bson:"linkMedidas1" json:"linkMedidas1,omitempty"`
	DescripcionMedidas2       string             `bson:"descripcionMedidas2" json:"descripcionMedidas2,omitempty"`
	LinkMedidas2              string             `bson:"linkMedidas2" json:"linkMedidas2,omitempty"`
	DescripcionMedidas3       string             `bson:"descripcionMedidas3" json:"descripcionMedidas3,omitempty"`
	LinkMedidas3              string             `bson:"linkMedidas3" json:"linkMedidas3,omitempty"`
	DescripcionMedidas4       string             `bson:"descripcionMedidas4" json:"descripcionMedidas4,omitempty"`
	LinkMedidas4              string             `bson:"linkMedidas4" json:"linkMedidas4,omitempty"`
	DescripcionMedidas5       string             `bson:"descripcionMedidas5" json:"descripcionMedidas5,omitempty"`
	LinkMedidas5              string             `bson:"linkMedidas5" json:"linkMedidas5,omitempty"`
	Palabra1                  string             `bson:"palabra1" json:"palabra1,omitempty"`
	Palabra2                  string             `bson:"palabra2" json:"palabra2,omitempty"`
	Palabra3                  string             `bson:"palabra3" json:"palabra3,omitempty"`
	Palabra4                  string             `bson:"palabra4" json:"palabra4,omitempty"`
	Palabra5                  string             `bson:"palabra5" json:"palabra5,omitempty"`
	Palabra6                  string             `bson:"palabra6" json:"palabra6,omitempty"`
	Palabra7                  string             `bson:"palabra7" json:"palabra7,omitempty"`
	Palabra8                  string             `bson:"palabra8" json:"palabra8,omitempty"`
	Palabra9                  string             `bson:"palabra9" json:"palabra9,omitempty"`
	Palabra10                 string             `bson:"palabra10" json:"palabra10,omitempty"`
	DetallePrensa             string             `bson:"detallePrensa" json:"detallePrensa,omitempty"`
	DescripcionPrensa1        string             `bson:"descripcionPrensa1" json:"descripcionPrensa1,omitempty"`
	LinkPrensa1               string             `bson:"linkPrensa1" json:"linkPrensa1,omitempty"`
	DescripcionPrensa2        string             `bson:"descripcionPrensa2" json:"descripcionPrensa2,omitempty"`
	LinkPrensa2               string             `bson:"linkPrensa2" json:"linkPrensa2,omitempty"`
	DescripcionPrensa3        string             `bson:"descripcionPrensa3" json:"descripcionPrensa3,omitempty"`
	LinkPrensa3               string             `bson:"linkPrensa3" json:"linkPrensa3,omitempty"`
	DescripcionPrensa4        string             `bson:"descripcionPrensa4" json:"descripcionPrensa4,omitempty"`
	LinkPrensa4               string             `bson:"linkPrensa4" json:"linkPrensa4,omitempty"`
	DescripcionPrensa5        string             `bson:"descripcionPrensa5" json:"descripcionPrensa5,omitempty"`
	LinkPrensa5               string             `bson:"linkPrensa5" json:"linkPrensa5,omitempty"`
	DetalleReaccionTerritorio string             `bson:"detalleReaccionTerritorio" json:"detalleReaccionTerritorio,omitempty"`
	DescripcionTerritorio     string             `bson:"descripcionTerritorio" json:"descripcionTerritorio,omitempty"`
	LinkTerritorio1           string             `bson:"linkTerritorio1" json:"linkTerritorio1,omitempty"`
	LinkTerritorio2           string             `bson:"linkTerritorio2" json:"linkTerritorio2,omitempty"`
	LinkTerritorio3           string             `bson:"linkTerritorio3" json:"linkTerritorio3,omitempty"`
	LinkTerritorio4           string             `bson:"linkTerritorio4" json:"linkTerritorio4,omitempty"`
	LinkTerritorio5           string             `bson:"linkTerritorio5" json:"linkTerritorio5,omitempty"`
	DetalleRelacionLabor      string             `bson:"detalleRelacionLabor" json:"detalleRelacionLabor,omitempty"`
	DescripcionRelacionLabor  string             `bson:"descripcionRelacionLabor" json:"descripcionRelacionLabor,omitempty"`
	LinkRelacionLabor1        string             `bson:"linkRelacionLabor1" json:"linkRelacionLabor1,omitempty"`
	LinkRelacionLabor2        string             `bson:"linkRelacionLabor2" json:"linkRelacionLabor2,omitempty"`
	LinkRelacionLabor3        string             `bson:"linkRelacionLabor3" json:"linkRelacionLabor3,omitempty"`
	LinkRelacionLabor4        string             `bson:"linkRelacionLabor4" json:"linkRelacionLabor4,omitempty"`
	LinkRelacionLabor5        string             `bson:"linkRelacionLabor5" json:"linkRelacionLabor5,omitempty"`
}
