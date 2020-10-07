import React from "react";
import { map, isEmpty } from "lodash";
import User from "./User";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faExclamationTriangle } from "@fortawesome/free-solid-svg-icons";

import "./ListaUsuarios.scss";

export default function (props) {
  const { users } = props;

  if (isEmpty(users)) {
    return (
      <div className="error-usuarios">
        <FontAwesomeIcon icon={faExclamationTriangle} />
        <h2>No se encuentran usuarios registrados</h2>
      </div>
    );
  }

  return (
    <ul className="lista-usuarios">
      {map(users, (user) => (
        <User key={user.id} user={user} />
      ))}
    </ul>
  );
}
