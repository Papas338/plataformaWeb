import React, { useState } from 'react';
import { Container, Row, Col, Form, Button} from "react-bootstrap";
import "./InfoLideres.scss";
import { Image, Page, Text, View, Document, StyleSheet, PDFDownloadLink } from '@react-pdf/renderer';
import { Camera } from "../../utils/icons";
import { descargaInfo } from "../../api/lideres"

function obtenerImagen(nombre) {    
    return require('../../assests/lideres/'+nombre+'.jpg')
}

export default function InfoLideres(props) {
    const {lider, imagenes} = props;

    var fileName = lider.nombre + ".pdf"

    // Create styles
    const styles = StyleSheet.create({
        page: {
        backgroundColor: '#ffffff'
        },
        section: {
        margin: 10,
        padding: 10,
        flexGrow: 1
        }
    });

    const [show, setShow] = useState(true);

    const descarga = () => {        
        descargaInfo(lider.id);
    }    
    
    // Create Document Component
    const MyDocument = () => (      
        <Document>
        <Page size="A4" style={styles.page}>
            <View style={styles.section}>
                <Text>Resumen</Text>
                <Text>Nombre: {lider.nombre}</Text>
                <Text>Fecha: {lider.fecha}</Text>
                <Text>Tipo liderazgo: {lider.tipoLiderazgo}</Text>
                <Text>Departamento: {lider.departamento}</Text>
                <Text>Municipio: {lider.municipio}</Text>
                <Text>Territorio: {lider.territorio}</Text>
            </View>
            <View style={styles.section}>
                <Text>Relacion Labor</Text>
                <Text>Detalle: {lider.detalleRelacionLabor}</Text>
                <Text>Descripción: {lider.descripcionRelacionLabor}</Text>
            </View>
            <View style={styles.section}>
                <Text>Prensa</Text>
                <Text>Detalle: {lider.detallePrensa}</Text>
                <Text>Descripción 1: {lider.descripcionPrensa1}</Text>
                <Text>Descripción 2: {lider.descripcionPrensa2}</Text>
                <Text>Descripción 3: {lider.descripcionPrensa3}</Text>
                <Text>Descripción 4: {lider.descripcionPrensa4}</Text>
                <Text>Descripción 5: {lider.descripcionPrensa5}</Text>
            </View>
            <View style={styles.section}>
                <Text>Medidas de reparación</Text>
                <Text>Detalle: {lider.detalleReparacion}</Text>
                <Text>Descripción 1: {lider.descripcionMedidas1}</Text>
                <Text>Descripción 2: {lider.descripcionMedidas2}</Text>
                <Text>Descripción 3: {lider.descripcionMedidas3}</Text>
                <Text>Descripción 4: {lider.descripcionMedidas4}</Text>
                <Text>Descripción 5: {lider.descripcionMedidas5}</Text>
            </View>
            <View style={styles.section}>
                <Text>Reacción territorio</Text>
                <Text>Detalle: {lider.detalleReaccionTerritorio}</Text>
                <Text>Descripción: {lider.descripcionTerritorio}</Text>
            </View>
            <View style={styles.section}>
                <Text>Palabras clave</Text>
                <Text>Palabra 1: {lider.palabra1}</Text>
                <Text>Palabra 2: {lider.palabra2}</Text>
                <Text>Palabra 3: {lider.palabra3}</Text>
                <Text>Palabra 4: {lider.palabra4}</Text>
                <Text>Palabra 5: {lider.palabra5}</Text>
                <Text>Palabra 6: {lider.palabra6}</Text>
                <Text>Palabra 7: {lider.palabra7}</Text>
                <Text>Palabra 8: {lider.palabra8}</Text>
                <Text>Palabra 9: {lider.palabra9}</Text>
                <Text>Palabra 10: {lider.palabra10}</Text>
            </View>
        </Page>
        </Document>
    );
    
    return (
        <Container className="info-lideres">
            <h1>Información Lider</h1>
            <Row>
                <Col sm={12}>
                    <fieldset className="formulario-lider__resumen">
                        <legend>Imagenes</legend>
                        <Row>
                            <Col sm={3}>
                                {!imagenes.imagen1 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.usuarioId+"1")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen2 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.usuarioId+"2")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen3 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.usuarioId+"3")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen4 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.usuarioId+"4")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                        </Row>
                        <Row>
                            <Col sm={3}>
                                {!imagenes.imagen5 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.usuarioId+"5")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen6 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.usuarioId+"6")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen7 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.usuarioId+"7")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen8 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.usuarioId+"8")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                        </Row>
                        <Row>
                            <Col sm={3}>
                                {!imagenes.imagen9 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.usuarioId+"9")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen10 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.usuarioId+"10")} alt="imagen1" height="100px" width="100px" />
                                )}  
                            </Col>
                        </Row>
                    </fieldset>
                </Col>
            </Row>
            <Row>
                <fieldset>
                    <legend>Resumen</legend>
                    <Row>
                        <Col>
                            <p>Nombre:</p>
                        </Col>
                        <Col>
                            {lider.nombre}
                        </Col>
                    </Row>
                    <Row>
                        <Col>
                            <p>Fecha:</p>
                        </Col>
                        <Col>
                            {lider.fecha}
                        </Col>
                    </Row>
                    <Row>
                        <Col>
                            <p>Municipio:</p>
                        </Col>
                        <Col>
                            {lider.municipio}
                        </Col>
                    </Row>
                    <Row>
                        <Col>
                            <p>Departamento:</p>
                        </Col>
                        <Col>
                            {lider.departamento}
                        </Col>
                    </Row>
                    <Row>
                        <Col>
                            <p>Tipo de liderazgo:</p>
                        </Col>
                        <Col>
                            {lider.tipoLiderazgo}
                        </Col>
                    </Row>
                    <Row>
                        <Col>
                            <p>Territorio:</p>
                        </Col>
                        <Col>
                            {lider.territorio}
                        </Col>
                    </Row>
                </fieldset>                                
            </Row>            
            <Row className="justify-content-md-center">                    
                <Button onClick={() => {
                    descarga();
                    setShow(!show);
                }} variant="success">
                   Descargar información
                </Button>
            </Row>
            {show ? (
                <Row className="justify-content-md-center">
                    Descargue un PDF con todos los datos del lider social
                </Row>
            ) : (                
                <Row className="justify-content-md-center">
                    Link de descarga:               
                    <PDFDownloadLink document={<MyDocument />} fileName={fileName} variant="success">
                        {({ blob, url, loading, error }) => (loading ? 'Cargando información...' : 'Información de lider')}
                    </PDFDownloadLink>
                </Row>
            )}
        </Container>
    )
}