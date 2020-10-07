import React from "react";
import { map, isEmpty } from "lodash";
import User from "./User";

import "./ListaUsuarios.scss";

export default function (props) {
  const { users } = props;

  if (isEmpty(users)) {
    return <h2>No se encuentran usuarios registrados</h2>;
  }

  return (
    <ul className="lista-usuarios">
      {map(users, (user) => (
        <User key={user.id} user={user} />
      ))}
    </ul>
  );
}
