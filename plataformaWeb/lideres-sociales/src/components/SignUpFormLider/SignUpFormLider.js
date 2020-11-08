import React, { useState } from "react";
import { Container, Row, Col, Form, Button, Spinner } from "react-bootstrap";
import { values, size } from "lodash";
import { toast } from "react-toastify";
import { signUpLiderTemp } from "../../api/lideres";
import { Link } from "react-router-dom";

import "./SignUpFormLider.scss"

export default function SignUpFormLider(props) {
    const [formData, setFormData] = useState(initialFormValue());
    const [signUpLoading, setSignUpLoading] = useState(false);

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
            setSignUpLoading(true);
            signUpLiderTemp(formData).then(response => {
                if(response.code) {
                    toast.warning(response.message);
                } else {
                    toast.success("El registro ha sido correcto");
                    setFormData(initialFormValue());
                }
            }).catch(() => {
                toast.error("Error del servidor, inténtelo más tarde!");
            }).finally(() => {
                setSignUpLoading(false);
            })
                
    };

    const onChange = e => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    }

    return (
        <Container fluid="md">
            <Form onSubmit={onSubmit} onChange={onChange}>
                <Form.Group>
                    <Row>
                        <Col sm={6}>
                            <fieldset className="formulario-lider__resumen">
                                <legend>Resumen</legend>
                                <Row>
                                    <Col>
                                        <h3>Nombre:</h3>
                                    </Col>
                                    <Col>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte nombre"
                                            name="nombre"
                                            defaultValue={formData.nombre}
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Col>
                                        <h3>Fecha:</h3>
                                    </Col>
                                    <Col>
                                        <Form.Control
                                            type="date"
                                            name="fecha"
                                            defaultValue={formData.fecha}
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Col>
                                        <h3>Municipio:</h3>
                                    </Col>
                                    <Col>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte municipio"
                                            name="municipio"
                                            defaultValue={formData.municipio}
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Col>
                                        <h3>Departamento:</h3>
                                    </Col>
                                    <Col>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte departamento"
                                            name="departamento"
                                            defaultValue={formData.departamento}
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Col>
                                        <h3>Tipo de liderazgo:</h3>
                                    </Col>
                                    <Col>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte tipo de liderazgo"
                                            name="tipoLiderazgo"
                                            defaultValue={formData.tipoLiderazgo}
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Col>
                                        <h3>Territorio:</h3>
                                    </Col>
                                    <Col>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte territorio"
                                            name="territorio"
                                            defaultValue={formData.territorio}
                                        />
                                    </Col>
                                </Row>
                            </fieldset>
                        </Col>
                        <Col sm={6}>
                            <fieldset className="formulario-lider__labor-relacion">
                                <legend>Labor/Relación</legend>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Detalle</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte detalle de labor/relación"
                                            name="detalleRelacionLabor"
                                            defaultValue={formData.detalleRelacionLabor}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de labor/relación"
                                            name="descripcionRelacionLabor"
                                            defaultValue={formData.descripcionRelacionLabor}
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Form.Row>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 1</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 1"
                                                name="linkRelacionLabor1"
                                                defaultValue={formData.linkRelacionLabor1} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 2</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 2"
                                                name="linkRelacionLabor2"
                                                defaultValue={formData.linkRelacionLabor2} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 3</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 3"
                                                name="linkRelacionLabor3"
                                                defaultValue={formData.linkRelacionLabor3} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 4</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 4"
                                                name="linkRelacionLabor4"
                                                defaultValue={formData.linkRelacionLabor4} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 5</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 5"
                                                name="linkRelacionLabor5"
                                                defaultValue={formData.linkRelacionLabor5} 
                                            />
                                        </Form.Group>
                                    </Form.Row>
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
                                        <Form.Label>Detalle</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte detalle de prensa"
                                            name="detallePrensa"
                                            defaultValue={formData.detallePrensa}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción 1</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de prensa"
                                            name="descripcionPrensa1"
                                            defaultValue={formData.descripcionPrensa1}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción 2</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de prensa"
                                            name="descripcionPrensa2"
                                            defaultValue={formData.descripcionPrensa2}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción 3</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de prensa"
                                            name="descripcionPrensa3"
                                            defaultValue={formData.descripcionPrensa3}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción 4</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de prensa"
                                            name="descripcionPrensa4"
                                            defaultValue={formData.descripcionPrensa4}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción 5</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de prensa"
                                            name="descripcionPrensa5"
                                            defaultValue={formData.descripcionPrensa5}
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Form.Row>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 1</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 1"
                                                name="linkPrensa1"
                                                defaultValue={formData.linkPrensa1} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 2</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 2"
                                                name="linkPrensa2"
                                                defaultValue={formData.linkPrensa2} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 3</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 3"
                                                name="linkPrensa3"
                                                defaultValue={formData.linkPrensa3} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 4</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 4"
                                                name="linkPrensa4"
                                                defaultValue={formData.linkPrensa4} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 5</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 5"
                                                name="linkPrensa5"
                                                defaultValue={formData.linkPrensa5} 
                                            />
                                        </Form.Group>
                                    </Form.Row>
                                </Row>
                            </fieldset>
                        </Col>
                        <Col sm={6}>
                            <fieldset className="formulario-lider__medidas-reparacion">
                                <legend>Medidas de reparación</legend>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Detalle</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte detalle de reparación"
                                            name="detallePrensa"
                                            defaultValue={formData.detalleReparacion}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción 1</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de medidas"
                                            name="descripcionPrensa1"
                                            defaultValue={formData.descripcionMedidas1}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción 2</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de medidas"
                                            name="descripcionPrensa2"
                                            defaultValue={formData.descripcionMedidas2}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción 3</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de medidas"
                                            name="descripcionPrensa3"
                                            defaultValue={formData.descripcionMedidas3}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción 4</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de medidas"
                                            name="descripcionPrensa4"
                                            defaultValue={formData.descripcionMedidas4}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción 5</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de medidas"
                                            name="descripcionPrensa5"
                                            defaultValue={formData.descripcionMedidas5}
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Form.Row>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 1</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 1"
                                                name="linkPrensa1"
                                                defaultValue={formData.linkMedidas1} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 2</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 2"
                                                name="linkPrensa2"
                                                defaultValue={formData.linkMedidas2} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 3</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 3"
                                                name="linkPrensa3"
                                                defaultValue={formData.linkMedidas3} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 4</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 4"
                                                name="linkPrensa4"
                                                defaultValue={formData.linkMedidas4} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 5</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 5"
                                                name="linkPrensa5"
                                                defaultValue={formData.linkMedidas5} 
                                            />
                                        </Form.Group>
                                    </Form.Row>
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
                                        <Form.Label>Detalle</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte detalle de intereses"
                                            name="detalleIntereses"
                                            defaultValue={formData.detalleIntereses}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de intereses"
                                            name="descripcionIntereses"
                                            defaultValue={formData.descripcionIntereses}
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Form.Row>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 1</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 1"
                                                name="linkIntereses1"
                                                defaultValue={formData.linkIntereses1} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 2</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 2"
                                                name="linkIntereses2"
                                                defaultValue={formData.linkIntereses2} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 3</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 3"
                                                name="linkIntereses3"
                                                defaultValue={formData.linkIntereses3} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 4</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 4"
                                                name="linkIntereses4"
                                                defaultValue={formData.linkIntereses4} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 5</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 5"
                                                name="linkIntereses5"
                                                defaultValue={formData.linkIntereses5} 
                                            />
                                        </Form.Group>
                                    </Form.Row>
                                </Row>
                            </fieldset>
                        </Col>
                        <Col sm={6}>
                            <fieldset className="formulario-lider__reaccion">
                                <legend>Reacción comunidad/territorio</legend>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Detalle</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte detalle de reacción comunidad/territorio"
                                            name="detalleReaccionTerritorio"
                                            defaultValue={formData.detalleReaccionTerritorio}
                                        />
                                    </Col>
                                </Row>
                                <Row>                                    
                                    <Col>
                                        <Form.Label>Descripción</Form.Label>
                                        <Form.Control
                                            as="textarea"
                                            rows={2}
                                            placeholder="Inserte descripción de reacción comunidad/territorio"
                                            name="descripcionTerritorio"
                                            defaultValue={formData.descripcionTerritorio}
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Form.Row>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 1</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 1"
                                                name="linkTerritorio1"
                                                defaultValue={formData.linkTerritorio1} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 2</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 2"
                                                name="linkTerritorio2"
                                                defaultValue={formData.linkTerritorio2} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 3</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 3"
                                                name="linkTerritorio3"
                                                defaultValue={formData.linkTerritorio3} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 4</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 4"
                                                name="linkTerritorio4"
                                                defaultValue={formData.linkTerritorio4} 
                                            />
                                        </Form.Group>
                                        <Form.Group as={Col} md="4">
                                            <Form.Label>Link 5</Form.Label>
                                            <Form.Control
                                                type="text"
                                                placeholder="Inserte Link 5"
                                                name="linkTerritorio5"
                                                defaultValue={formData.linkTerritorio5} 
                                            />
                                        </Form.Group>
                                    </Form.Row>
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
                                        <h3>Palabra 1:</h3>
                                    </Col>
                                    <Col md={3}>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte palabra 1"
                                            name="linkpalabra1"
                                            defaultValue={formData.palabra1} 
                                        />
                                    </Col>
                                    <Col md={3}>
                                        <h3>Palabra 6:</h3>
                                    </Col>
                                    <Col md={3}>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte palabra 6"
                                            name="linkpalabra6"
                                            defaultValue={formData.palabra6} 
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Col md={3}>
                                        <h3>Palabra 2:</h3>
                                    </Col>
                                    <Col md={3}>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte palabra 2"
                                            name="linkpalabra2"
                                            defaultValue={formData.palabra2} 
                                        />
                                    </Col>
                                    <Col md={3}>
                                        <h3>Palabra 7:</h3>
                                    </Col>
                                    <Col md={3}>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte palabra 7"
                                            name="linkpalabra7"
                                            defaultValue={formData.palabra7} 
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Col md={3}>
                                        <h3>Palabra 3:</h3>
                                    </Col>
                                    <Col md={3}>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte palabra 3"
                                            name="linkpalabra3"
                                            defaultValue={formData.palabra3} 
                                        />
                                    </Col>
                                    <Col md={3}>
                                        <h3>Palabra 8:</h3>
                                    </Col>
                                    <Col md={3}>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte palabra 8"
                                            name="linkpalabra8"
                                            defaultValue={formData.palabra8} 
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Col md={3}>
                                        <h3>Palabra 4:</h3>
                                    </Col>
                                    <Col md={3}>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte palabra 4"
                                            name="linkpalabra4"
                                            defaultValue={formData.palabra4} 
                                        />
                                    </Col>
                                    <Col md={3}>
                                        <h3>Palabra 9:</h3>
                                    </Col>
                                    <Col md={3}>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte palabra 9"
                                            name="linkpalabra9"
                                            defaultValue={formData.palabra9} 
                                        />
                                    </Col>
                                </Row>
                                <Row>
                                    <Col md={3}>
                                        <h3>Palabra 5:</h3>
                                    </Col>
                                    <Col md={3}>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte palabra 5"
                                            name="linkpalabra5"
                                            defaultValue={formData.palabra5} 
                                        />
                                    </Col>
                                    <Col md={3}>
                                        <h3>Palabra 10:</h3>
                                    </Col>
                                    <Col md={3}>
                                        <Form.Control
                                            type="text"
                                            placeholder="Inserte palabra 10"
                                            name="linkpalabra10"
                                            defaultValue={formData.palabra10} 
                                        />
                                    </Col>
                                </Row>
                            </fieldset>
                        </Col>
                    </Row>
                </Form.Group>
                <Row className="justify-content-md-center">
                    <Col md="auto">
                        <Button variant="success" type="submit">
                            {!signUpLoading ? "Registrar Lider" : <Spinner animation="border" />}
                        </Button>
                    </Col>
                    <Col md="auto">
                        <Button variant="success" as={Link} to="/">Volver al inicio</Button>
                    </Col>
                </Row>
            </Form>
        </Container>
    );
}

function initialFormValue() {
    return {
        nombre: "",
        fecha: "",
        departamento: "",
        municipio: "",        
        tipoLiderazgo: "",
        territorio: "",
        usuarioId: "",
        detalleIntereses: "",
        descripcionIntereses: "",
        linkIntereses1: "",
        linkIntereses2: "",
        linkIntereses3: "",
        linkIntereses4: "",
        linkIntereses5: "",
        detalleReparacion: "",
        descripcionMedidas1: "",
        LinkMedidas1: "",
        descripcionMedidas2: "",
        LinkMedidas2: "",
        descripcionMedidas3: "",
        LinkMedidas3: "",
        descripcionMedidas4: "",
        LinkMedidas4: "",
        descripcionMedidas5: "",
        LinkMedidas5: "",
        palabra1: "",
        palabra2: "",
        palabra3: "",
        palabra4: "",
        palabra5: "",
        palabra6: "",
        palabra7: "",
        palabra8: "",
        palabra9: "",
        palabra10: "",
        detallePrensa: "",
        descripcionPrensa1: "",
        linkPrensa1: "",
        descripcionPrensa2: "",
        linkPrensa2: "",
        descripcionPrensa3: "",
        linkPrensa3: "",
        descripcionPrensa4: "",
        linkPrensa4: "",
        descripcionPrensa5: "",
        linkPrensa5: "",
        detalleReaccionTerritorio: "",
        descripcionTerritorio: "",
        linkTerritorio1: "",
        linkTerritorio2: "",
        linkTerritorio3: "",
        linkTerritorio4: "",
        linkTerritorio5: "",
        detalleRelacionLabor: "",
        descripcionRelacionLabor: "",
        linkRelacionLabor1: "",
        linkRelacionLabor2: "",
        linkRelacionLabor3: "",
        linkRelacionLabor4: "",
        linkRelacionLabor5: "",
    };
}