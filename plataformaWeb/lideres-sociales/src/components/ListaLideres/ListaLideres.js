import React from "react";
import { map, isEmpty } from "lodash";
import Lider from "./Lider";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faExclamationTriangle } from "@fortawesome/free-solid-svg-icons";

import "./ListaLideres.scss";

export default function (props) {
  const { openModal, setShowModal, lideres } = props;  

  console.log(props)

  if (isEmpty(lideres)) {   
      return (
        <div className="error-usuarios">
          <FontAwesomeIcon icon={faExclamationTriangle} />
          <h2>No se encuentran usuarios registrados</h2>
        </div>
      );       
  }

  return (
    <ul className="lista-usuarios">
      {map(lideres, (lider) => (
        <Lider openModal={openModal} setShowModal={setShowModal} key={lider.id} lider={lider} />
      ))}
    </ul>
  );
}