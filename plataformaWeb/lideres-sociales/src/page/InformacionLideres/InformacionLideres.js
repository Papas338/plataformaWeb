import React, { useState, useEffect } from "react";
import { Spinner, Container, Col, Row, Button, InputGroup } from "react-bootstrap";
import { Link, withRouter } from "react-router-dom";
import { isEmpty } from "lodash";
import queryString from "query-string";
import ListaLideres from "../../components/ListaLideres";
import { logoutApi } from "../../api/auth"
import { getLideres } from "../../api/lideres";
import InfoModal from "../../components/Modal/InfoModal";
import "./InformacionLideres.scss";

function InformacionLideres(props) {
  const { setRefreshCheckLogin, location, history } = props;
  const [lideres, setLideres] = useState(null);
  const params = useUsersQuery(location);
  const [btnLoading, setBtnLoading] = useState(false)
  const [showModal, setShowModal] = useState(false);
  const [contentModal, setContentModal] = useState(null);
  const openModal = (content) => {
    setShowModal(true);
    setContentModal(content);
  };

  const logout = () => {
    logoutApi();
    setRefreshCheckLogin(true);
  }

  useEffect(() => {
    getLideres(queryString.stringify(params))
      .then((response) => {
        if (params.page == 1) {
          if (isEmpty(response)) {
            setLideres([]);
          } else {
            setLideres(response);
          }
        } else {
          if (!response) {
            setBtnLoading(0);
          } else {
            setLideres([...lideres, ...response]);
            setBtnLoading(false);
          }
        }
      })
      .catch(() => {
        setLideres([]);
      });
  }, [location]);

  const moreData = () => {
    setBtnLoading(true);
    const newPage = parseInt(params.page) + 1;
    history.push({
      search: queryString.stringify({ ...params, page: newPage }),
    });
  }

  return (
    <>
      <Container>
        <Row>
          <Col className="users__header" sm={8}>
            <h2>Listado de lideres</h2>
          </Col>
          <Col className="users__headerLogout" sm={4}>
            <Button as={Link} variant="danger" onClick={logout}>Cerrar Sesión</Button>
          </Col>
        </Row>
        <Row>
          <Col className="users__search" sm={4}>
            <input type="text" placeholder="Buscar usuario"
              onChange={(e) => history.push({
                search: queryString.stringify({ ...params, search: e.target.value, page: 1 })
              })}
            />
            <Button variant="success" as={Link} to="/">
              Volver al inicio
            </Button>
            <p>
              Este módulo permite consultar los lideres indigenas asesinados.
            </p>          
          </Col>
          <Col className="users__usuarios" sm={8}>
            {!lideres ? (
              <div className="users__loading">
                <Spinner animation="border" variant="info" />
              Buscando lideres
              </div>
            ) : (
                <>
                  <ListaLideres openModal={openModal} setShowModal={setShowModal} lideres={lideres} />
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
      <InfoModal show={showModal} setShow={setShowModal}>
        {contentModal}
      </InfoModal>
    </>
  );
}

function useUsersQuery(location) {
  const { page = 1, search } = queryString.parse(location.search);

  return { page, search };
}

export default withRouter(InformacionLideres);
