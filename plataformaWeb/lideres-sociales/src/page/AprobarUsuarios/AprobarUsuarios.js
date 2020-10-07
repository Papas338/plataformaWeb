import React, { useState, useEffect } from "react";
import { Spinner, ButtonGroup, Button } from "react-bootstrap";
import { withRouter } from "react-router-dom";
import { isEmpty } from "lodash";
import queryString from "query-string";
import ListaUsuarios from "../../components/ListaUsuarios";
import { getUsuarios } from "../../api/usuarios";

import "./AprobarUsuarios.scss";

function AprobarUsuarios(props) {
  const { location } = props;
  const [users, setUsers] = useState(null);
  const params = useUsersQuery(location);

  //console.log(params);

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
    <div>
      <div className="users__title">
        <h2>Usuarios registrados</h2>
        <input type="text" placeholder="Buscar usuario" />
      </div>

      {!users ? (
        <div className="users__loading">
          <Spinner animation="border" variant="info" />
          Buscando usuarios
        </div>
      ) : (
        <ListaUsuarios users={users} />
      )}
    </div>
  );
}

function useUsersQuery(location) {
  const { page = 1, search } = queryString.parse(location.search);

  return { page, search };
}

export default withRouter(AprobarUsuarios);
