import React from "react";
import { Row, Col, Button } from "react-bootstrap";
import InfoLideres from "../../components/InfoLideres/InfoLideres";

export default function Lider(props) {
  const { openModal, setShowModal, lider } = props;
  console.log(lider);

    return (
      <Row className="lista-usuarios__user">
        <Col id="usuario" sm={8}>
            <h2>{lider.nombre}</h2>
            <h3>{lider.departamento} - {lider.municipio} - {lider.territorio}</h3>
        </Col>
        <Col id="boton" sm={4}>
          <Button variant="success" 
            onClick={() => openModal(<
              InfoLideres lider={lider} setShowModal={setShowModal}/>)
            }>
            Ver información
          </Button>
        </Col>
      </Row>
    );
}