import React from "react";
import { map, isEmpty } from "lodash";
import User from "./User";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faExclamationTriangle } from "@fortawesome/free-solid-svg-icons";

import "./ListaUsuarios.scss";

export default function (props) {
  const { users } = props;

  console.log(props)

  if (isEmpty(users)) {
    if (users.role == "Administrador") {
      return (
        <div className="error-usuarios">
          <FontAwesomeIcon icon={faExclamationTriangle} />
          <h2>No se encuentran usuarios registrados</h2>
        </div>
      );
    } else if ((users.role == "Usuario") || (users.role == "Pasante")) {
      return (
        <div className="error-usuarios">
          <FontAwesomeIcon icon={faExclamationTriangle} />
          <h2>No tiene permisos de administrador</h2>
        </div>
      );
    } else {
      return (
        <div className="error-usuarios">
          <FontAwesomeIcon icon={faExclamationTriangle} />
          <h2>Usuario(s) no encontrado(s)</h2>
        </div>
      );
    }    
  }

  return (
    <ul className="lista-usuarios">
      {map(users, (user) => (
        <User key={user.id} user={user} />
      ))}
    </ul>
  );
}
