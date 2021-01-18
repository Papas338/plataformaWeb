import React from 'react';
import { Spinner, Container, Col, Row, Button, InputGroup } from "react-bootstrap";
import { logoutApi } from "../../api/auth";
import { Link, withRouter } from "react-router-dom";

import "./VisualizarLideres.scss";

export default function VisualizarLideres(props) {

    const { setRefreshCheckLogin } = props;

    const logout = () => {
        logoutApi();
        setRefreshCheckLogin(true);
    }

    return (
        <>
            <Row>
                <Col className="visuaLider__header" sm={8}>
                    <h2>Ver Líderes</h2>
                </Col>
                <Col className="visuaLider__headerIndex" sm={2}>
                    <Button variant="success" as={Link} to="/">Volver al inicio</Button>
                </Col>
                <Col className="visuaLider__headerLogout" sm={2}>
                    <Button as={Link} variant="danger" onClick={logout}>Cerrar Sesión</Button>
                </Col>
            </Row>
            <Row>
                <Col>
                    <iframe width="100%" height="600px" src="https://udistritalfjc.maps.arcgis.com/apps/MapJournal/index.html?appid=52ce2fbe3e664b1bb2864b42e87aa340"></iframe>
                </Col>
            </Row>   
        </>
        
    )
}
