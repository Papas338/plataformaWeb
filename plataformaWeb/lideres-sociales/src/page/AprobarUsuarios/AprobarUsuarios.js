import React, { useState, useEffect } from "react";
import { Spinner, ButtonGroup, Button } from "react-bootstrap";
import { withRouter } from "react-router-dom";
import queryString from "query-string";
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
        console.log(response);
      })
      .catch(() => {
        setUsers([]);
      });
  }, []);

  return (
    <div className="users__title">
      <h2>Usuarios registrados</h2>
      <input type="text" placeholder="Buscar usuario" />
    </div>
  );
}

function useUsersQuery(location) {
  const { page = 1, search } = queryString.parse(location.search);

  return { page, search };
}

export default withRouter(AprobarUsuarios);
