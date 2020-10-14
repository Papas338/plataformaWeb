import React, { useState, useEffect } from "react";
import { Spinner, Container, Col, Row, Button, InputGroup } from "react-bootstrap";
import { Link, withRouter } from "react-router-dom";
import { isEmpty } from "lodash";
import queryString from "query-string";
import ListaUsuarios from "../../components/ListaUsuarios";
import { logoutApi } from "../../api/auth"
import { getUsuarios } from "../../api/usuarios";
import "./AprobarUsuarios.scss";

function AprobarUsuarios(props) {
  const { setRefreshCheckLogin, location, history } = props;
  const [users, setUsers] = useState(null);
  const params = useUsersQuery(location);
  const [btnLoading, setBtnLoading] = useState(false)

  const logout = () => {
    logoutApi();
    setRefreshCheckLogin(true);
  }

  useEffect(() => {
    getUsuarios(queryString.stringify(params))
      .then((response) => {
        if (params.page == 1)  {
          if (isEmpty(response)) {
            setUsers([]);
          } else {
            setUsers(response);
          }
        } else {
          if (!response) {
            setBtnLoading(0);
          } else {
            setUsers([...users, ...response]);
            setBtnLoading(false);
          }
        }        
      })
      .catch(() => {
        setUsers([]);
      });
  }, [location]);
  
  const moreData = () => {
    setBtnLoading(true);
    const newPage = parseInt(params.page) + 1;
    history.push({
      search: queryString.stringify({...params, page: newPage}),
    });
  }

  return (
    <Container>
      <Row>
        <Col className="users__header" sm={8}>
          <h2>Listado de usuarios</h2>
        </Col>
        <Col className="users__headerLogout" sm={4}>
          <Button as={Link} variant="danger" onClick={logout}>Cerrar Sesión</Button>
        </Col>
      </Row>
      <Row>
        <Col className="users__search" sm={4}>
            <input type="text" placeholder="Buscar usuario" 
              onChange={(e) => history.push({
                search: queryString.stringify({...params, search: e.target.value, page:1}) 
              })} 
            />            
          <Button variant="success" as={Link} to="/">
            Volver al inicio
          </Button>
        </Col>
        <Col className="users__usuarios" sm={8}>
        {!users ? (
          <div className="users__loading">
            <Spinner animation="border" variant="info" />
            Buscando usuarios
          </div>
        ) : (
          <>
            <ListaUsuarios users={users} />
            <Button onClick={moreData} className="load-more">
              {!btnLoading ? (
                btnLoading !== 0 && "Cargar más usuarios"
              ) : (
                <Spinner                
                  as="span"
                  animation="grow"
                  size="sm"
                  role="status"
                  arie-hidden="true"
                />
              )}
            </Button>
          </>
        )}
        </Col>
      </Row>
    </Container>
  );
}

function useUsersQuery(location) {
  const { page = 1, search } = queryString.parse(location.search);

  return { page, search };
}

export default withRouter(AprobarUsuarios);
