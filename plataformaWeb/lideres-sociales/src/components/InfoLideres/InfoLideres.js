import React from 'react';
import { Container, Row, Col, Form, Button} from "react-bootstrap";
import "./InfoLideres.scss";

export default function InfoLideres(props) {
    const {lider} = props;
    
    return (
        <Container className="info-lideres">
            <h1>Información Lider</h1>
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
                <Button variant="success">
                    Descargar información
                </Button>
            </Row>
        </Container>
    )
}