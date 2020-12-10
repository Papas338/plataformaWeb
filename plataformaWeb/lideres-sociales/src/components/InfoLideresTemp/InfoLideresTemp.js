import React from 'react';
import { Container, Row, Col, Form, Button} from "react-bootstrap";
import { signUpLider, deleteLiderTemp, getLideres } from "../../api/lideres";
import { toast } from "react-toastify";
import "./InfoLideresTemp.scss";
import { Camera } from "../../utils/icons";

function obtenerImagen(nombre) {    
    return require('../../assests/lideres/'+nombre+'.jpg')
}

export default function InfoLideresTemp(props) {
    const {lider, setShowModal, imagenes} = props;

    const eliminarLiderTemp = () => {
        deleteLiderTemp(lider.id);
        setShowModal(false); 
        window.location.reload(); 
      }
    
        const onSubmit = e => {
            e.preventDefault();
            /* let validCount = 0;
            values(formData).some(value => {
                value && validCount++
                return null
            });
    
            if(validCount !== size(formData)) {
                toast.warning("Completa todo los campos del formulario");
            } else { */
                signUpLider(lider).then(response => {
                    if(response.code) {
                        toast.warning(response.message);
                    } else {
                        toast.success("El registro ha sido correcto");
                    }
                }).catch(() => {
                    toast.error("Error del servidor, inténtelo más tarde!");
                })                    
        };

        console.log("prueba")

    return (        
        <Container className="info-lideresTemp">
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
                                    <img src={obtenerImagen(lider.id+"1")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen2 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.id+"2")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen3 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.id+"3")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen4 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.id+"4")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                        </Row>
                        <Row>
                            <Col sm={3}>
                                {!imagenes.imagen5 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.id+"5")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen6 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.id+"6")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen7 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.id+"7")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen8 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.id+"8")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                        </Row>
                        <Row>
                            <Col sm={3}>
                                {!imagenes.imagen9 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.id+"9")} alt="imagen1" height="100px" width="100px" />
                                )}
                            </Col>
                            <Col sm={3}>
                                {!imagenes.imagen10 ? (
                                    <Camera />
                                ) : (
                                    <img src={obtenerImagen(lider.id+"10")} alt="imagen1" height="100px" width="100px" />
                                )}  
                            </Col>
                        </Row>
                    </fieldset>
                </Col>
            </Row>
            <Row>
                <Col sm={6}>
                    <fieldset className="formulario-lider__resumen">
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
                </Col>
                <Col sm={6}>
                    <fieldset className="formulario-lider__labor-relacion">
                        <legend>Labor/Relación</legend>
                        <Row>                                    
                            <Col>
                                <p>Detalle:</p>
                            </Col>
                            <Col>
                                {lider.detalleRelacionLabor}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripción:</p>
                            </Col>
                            <Col>
                                {lider.descripcionRelacionLabor}
                            </Col>
                        </Row>
                        <Row>
                            <Col>
                                <p>Link 1:</p>
                            </Col>
                            <Col>
                                {lider.linkRelacionLabor1}
                            </Col>
                        </Row>
                        <Row>
                            <Col>
                                <p>Link 2:</p>
                            </Col>
                            <Col>
                                {lider.linkRelacionLabor2}
                            </Col>
                        </Row>
                        <Row>
                            <Col>
                                <p>Link 3:</p>
                            </Col>
                            <Col>
                                {lider.linkRelacionLabor3}
                            </Col>
                        </Row>
                        <Row>
                            <Col>
                                <p>Link 4:</p>
                            </Col>
                            <Col>
                                {lider.linkRelacionLabor4}
                            </Col>
                        </Row>
                        <Row>
                            <Col>
                                <p>Link 5:</p>
                            </Col>
                            <Col>
                                {lider.linkRelacionLabor5}
                            </Col>
                        </Row>
                    </fieldset>
                </Col>
            </Row>
            <Row>
                <Col sm={6}>
                    <fieldset className="formulario-lider__prensa">
                        <legend>Prensa</legend>
                        <Row>                                    
                            <Col>
                                <p>Detalle:</p>
                            </Col>
                            <Col>
                                {lider.detallePrensa}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripción 1:</p>
                            </Col>
                            <Col>
                                {lider.descripcionPrensa1}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripción 2:</p>
                            </Col>
                            <Col>
                                {lider.descripcionPrensa2}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripción 3:</p>
                            </Col>
                            <Col>
                                {lider.descripcionPrensa3}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripción 4:</p>
                            </Col>
                            <Col>
                                {lider.descripcionPrensa4}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripción 5:</p>
                            </Col>
                            <Col>
                                {lider.descripcionPrensa5}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 1:</p>
                            </Col>
                            <Col>
                                {lider.linkPrensa1}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 2:</p>
                            </Col>
                            <Col>
                                {lider.linkPrensa2}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 3:</p>
                            </Col>
                            <Col>
                                {lider.linkPrensa3}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 4:</p>
                            </Col>
                            <Col>
                                {lider.linkPrensa4}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 5:</p>
                            </Col>
                            <Col>
                                {lider.linkPrensa5}
                            </Col>
                        </Row>
                    </fieldset>
                </Col>
                <Col sm={6}>
                    <fieldset className="formulario-lider__medidas-reparacion">
                        <legend>Medidas de reparación</legend>
                        <Row>                                    
                            <Col>
                                <p>Detalle:</p>
                            </Col>
                            <Col>
                                {lider.detalleReparacion}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripción 1:</p>
                            </Col>
                            <Col>
                                {lider.descripcionMedidas1}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripción 2:</p>
                            </Col>
                            <Col>
                                {lider.descripcionMedidas2}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripción 3:</p>
                            </Col>
                            <Col>
                                {lider.descripcionMedidas3}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripción 4:</p>
                            </Col>
                            <Col>
                                {lider.descripcionMedidas4}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripción 5:</p>
                            </Col>
                            <Col>
                                {lider.descripcionMedidas5}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 1:</p>
                            </Col>
                            <Col>
                                {lider.linkMedidas1}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 2:</p>
                            </Col>
                            <Col>
                                {lider.linkMedidas2}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 3:</p>
                            </Col>
                            <Col>
                                {lider.linkMedidas3}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 4:</p>
                            </Col>
                            <Col>
                                {lider.linkMedidas4}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 5:</p>
                            </Col>
                            <Col>
                                {lider.linkMedidas5}
                            </Col>
                        </Row>
                    </fieldset>
                </Col>
            </Row>
            <Row>
                <Col sm={6}>
                    <fieldset className="formulario-lider__intereses">
                        <legend>Intereses</legend>
                        <Row>                                    
                            <Col>
                                <p>Detalle:</p>
                            </Col>
                            <Col>
                                {lider.detalleIntereses}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripcion:</p>
                            </Col>
                            <Col>
                                {lider.descripcionIntereses}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 1:</p>
                            </Col>
                            <Col>
                                {lider.linkIntereses1}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 2:</p>
                            </Col>
                            <Col>
                                {lider.linkIntereses2}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 3:</p>
                            </Col>
                            <Col>
                                {lider.linkIntereses3}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 4:</p>
                            </Col>
                            <Col>
                                {lider.linkIntereses4}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 5:</p>
                            </Col>
                            <Col>
                                {lider.linkIntereses5}
                            </Col>
                        </Row>
                    </fieldset>
                </Col>
                <Col sm={6}>
                    <fieldset className="formulario-lider__reaccion">
                        <legend>Reacción comunidad/territorio</legend>
                        <Row>                                    
                            <Col>
                                <p>Detalle:</p>
                            </Col>
                            <Col>
                                {lider.detalleReaccionTerritorio}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Descripción:</p>
                            </Col>
                            <Col>
                                {lider.descripcionTerritorio}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 1:</p>
                            </Col>
                            <Col>
                                {lider.linkTerritorio1}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 2:</p>
                            </Col>
                            <Col>
                                {lider.linkTerritorio2}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 3:</p>
                            </Col>
                            <Col>
                                {lider.linkTerritorio3}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 4:</p>
                            </Col>
                            <Col>
                                {lider.linkTerritorio4}
                            </Col>
                        </Row>
                        <Row>                                    
                            <Col>
                                <p>Link 5:</p>
                            </Col>
                            <Col>
                                {lider.linkTerritorio5}
                            </Col>
                        </Row>
                    </fieldset>
                </Col>
            </Row>
            <Row>
                <Col sm={12}>
                    <fieldset className="formulario-lider__palabras">
                        <legend>Palabras claves</legend>
                        <Row>
                            <Col md={3}>
                                <p>Palabra 1:</p>
                            </Col>
                            <Col md={3}>
                                {lider.palabra1}
                            </Col>
                            <Col md={3}>
                                <p>Palabra 6:</p>
                            </Col>
                            <Col md={3}>
                                {lider.palabra6}
                            </Col>
                        </Row>
                        <Row>
                            <Col md={3}>
                                <p>Palabra 2:</p>
                            </Col>
                            <Col md={3}>
                                {lider.palabra2}
                            </Col>
                            <Col md={3}>
                                <p>Palabra 7:</p>
                            </Col>
                            <Col md={3}>
                                {lider.palabra7}
                            </Col>
                        </Row>
                        <Row>
                            <Col md={3}>
                                <p>Palabra 3:</p>
                            </Col>
                            <Col md={3}>
                                {lider.palabra3}
                            </Col>
                            <Col md={3}>
                                <p>Palabra 8:</p>
                            </Col>
                            <Col md={3}>
                                {lider.palabra8}
                            </Col>
                        </Row>
                        <Row>
                            <Col md={3}>
                                <p>Palabra 4:</p>
                            </Col>
                            <Col md={3}>
                                {lider.palabra4}
                            </Col>
                            <Col md={3}>
                                <p>Palabra 9:</p>
                            </Col>
                            <Col md={3}>
                                {lider.palabra9}
                            </Col>
                        </Row>
                        <Row>
                            <Col md={3}>
                                <p>Palabra 5:</p>
                            </Col>
                            <Col md={3}>
                                {lider.palabra5}
                            </Col>
                            <Col md={3}>
                                <p>Palabra 10:</p>
                            </Col>
                            <Col md={3}>
                                {lider.palabra10}
                            </Col>
                        </Row>
                    </fieldset>
                </Col> 
            </Row>
            <Row className="justify-content-md-center">
                    <Col md="auto">
                        <Form onSubmit={onSubmit}>
                            <Button variant="success" type="submit" onClick={() => eliminarLiderTemp()}>
                                Admitir
                            </Button>
                        </Form>
                    </Col>
                </Row>
        </Container>
    )
}
