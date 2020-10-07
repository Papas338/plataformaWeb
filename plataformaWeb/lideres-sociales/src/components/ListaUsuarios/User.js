import React from "react";
import { Row, Col, Button } from "react-bootstrap";

export default function User(props) {
  const { user } = props;

  return (
    <Row className="lista-usuarios__user">
      <Col id="usuario" sm={8}>
        <h2>{user.email}</h2>
        <h3>{user.role}</h3>
      </Col>
      <Col id="boton" sm={4}>
        <Button>Convertir</Button>
      </Col>
    </Row>
  );
}
