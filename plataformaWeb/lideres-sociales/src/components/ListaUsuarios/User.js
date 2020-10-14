import React from "react";
import { Row, Col, Button } from "react-bootstrap";
import { modificarUsuario } from "../../api/usuarios";

export default function User(props) {
  const { user } = props;

  const modificoUsuario = () => {
    modificarUsuario(user.id); 
    window.location.reload(); 
  }
  if (user.role != "Administrador") {
    return (
      <Row className="lista-usuarios__user">
        <Col id="usuario" sm={8}>
          <h2>{user.email}</h2>
          <h3>{user.role}</h3>
        </Col>
        <Col id="boton" sm={4}>
          <Button variant="success" onClick={() => modificoUsuario()}>Convertir</Button>
        </Col>
      </Row>
    );
  } else {
    return (
      <Row className="lista-usuarios__user">
        <Col id="usuario" sm={8}>
          <h2>{user.email}</h2>
          <h3>{user.role}</h3>
        </Col>
      </Row>
    );
  }
}
