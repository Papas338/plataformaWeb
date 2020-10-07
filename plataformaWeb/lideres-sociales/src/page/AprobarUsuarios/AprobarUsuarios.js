import React, { useState, useEffect } from "react";
import { Spinner, Container, Col, Row, Button } from "react-bootstrap";
import { Link, withRouter } from "react-router-dom";
import { isEmpty } from "lodash";
import queryString from "query-string";
import ListaUsuarios from "../../components/ListaUsuarios";
import { getUsuarios } from "../../api/usuarios";

import "./AprobarUsuarios.scss";

function AprobarUsuarios(props) {
  const { location } = props;
  const [users, setUsers] = useState(null);
  const params = useUsersQuery(location);

  useEffect(() => {
    getUsuarios(queryString.stringify(params))
      .then((response) => {
        if (isEmpty(response)) {
          setUsers([]);
        } else {
          setUsers(response);
        }
      })
      .catch(() => {
        setUsers([]);
      });
  }, []);

  return (
    <Container>
      <Row>
        <Header />
      </Row>
      <Row>
        <Search />
        <Usuarios users={users} />
      </Row>
    </Container>
  );
}

function Header() {
  return (
    <Col className="users__header">
      <h2>Listado de usuarios</h2>
    </Col>
  );
}

function Search() {
  return (
    <Col className="users__search" sm={4}>
      <input type="text" placeholder="Buscar usuario" />
      <Button as={Link} to="/">
        Volver al inicio
      </Button>
    </Col>
  );
}

function Usuarios(props) {
  const { users } = props;
  return (
    <Col className="users__usuarios" sm={8}>
      {!users ? (
        <div className="users__loading">
          <Spinner animation="border" variant="info" />
          Buscando usuarios
        </div>
      ) : (
        <ListaUsuarios users={users} />
      )}
    </Col>
  );
}

function useUsersQuery(location) {
  const { page = 1, search } = queryString.parse(location.search);

  return { page, search };
}

export default withRouter(AprobarUsuarios);
