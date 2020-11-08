import React from "react";
import { map, isEmpty } from "lodash";
import LiderTemp from "./LiderTemp";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faExclamationTriangle } from "@fortawesome/free-solid-svg-icons";

import "./ListaLideresTemp.scss";

export default function (props) {
  const { openModal, setShowModal, lideres } = props;

  if (isEmpty(lideres)) {   
      return (
        <div className="error-lideresTemp">
          <FontAwesomeIcon icon={faExclamationTriangle} />
          <h2>No se encuentran usuarios registrados</h2>
        </div>
      );       
  }

  return (
    <ul className="lista-lideresTemp">
      {map(lideres, (lider) => (
        <LiderTemp openModal={openModal} setShowModal={setShowModal} key={lider.id} lider={lider} />
      ))}
    </ul>
  );
}