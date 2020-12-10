import React from 'react';
import { Container, Row, Col, Button } from "react-bootstrap";
import { Link } from "react-router-dom";
import { logoutApi } from "../../api/auth";
import investigacion from "../../document/Documento final proyecto de investigacion Lideres Sociales.pdf"

export default function Investigacion(props) {

    const { setRefreshCheckLogin } = props;

    const logout = () => {
        logoutApi();
        setRefreshCheckLogin(true);
    }

    return (
        <>
            <Row>
                <Col className="visuaLider__header" sm={8}>
                    <h2>Ver Investigación</h2>
                </Col>
                <Col className="visuaLider__headerIndex" sm={2}>
                    <Button variant="success" as={Link} to="/">Volver al inicio</Button>
                </Col>
                <Col className="visuaLider__headerLogout" sm={2}>
                    <Button as={Link} variant="danger" onClick={logout}>Cerrar Sesión</Button>
                </Col>
            </Row>
            <Row>
                <embed src={investigacion} type="application/pdf" width="100%" height="600px" />
            </Row>
        </>
    )
}
