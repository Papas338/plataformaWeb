import React, { useState, useEffect } from "react";
import { Row, Col, Button } from "react-bootstrap";
import InfoLideres from "../../components/InfoLideres/InfoLideres";
import { getImagenes } from "../../api/lideres";
import { isEmpty } from "lodash";

export default function Lider(props) {
  const { openModal, setShowModal, lider } = props;
  
  const [imagenes, setImagenes] = useState(null);

  var cargaImagen = 0;

    console.log(lider.id)

    useEffect(() => {
      getImagenes(lider.id).then((response) => {
          if(isEmpty(response)) {
              setImagenes([]);            
          } else {
              setImagenes(response);
          }
      }).catch(() => {
          setImagenes([]);
        });
      console.log("effect")
      cargaImagen++;
    }, [cargaImagen])

    return (
      <Row className="lista-usuarios__user">
        <Col id="usuario" sm={8}>
            <h2>{lider.nombre}</h2>
            <h3>{lider.departamento} - {lider.municipio} - {lider.territorio}</h3>
        </Col>
        <Col id="boton" sm={4}>
          <Button variant="success" 
            onClick={() => openModal(<
              InfoLideres imagenes={imagenes} lider={lider} setShowModal={setShowModal}/>)
            }>
            Ver información
          </Button>
        </Col>
      </Row>
    );
}