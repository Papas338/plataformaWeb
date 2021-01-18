import React, {useState, useEffect} from "react";
import { Row, Col, Button, Spinner, Form } from "react-bootstrap";
import { signUpLider, signUpLiderEsri, deleteLiderTemp } from "../../api/lideres";
import { toast } from "react-toastify";
import { isEmpty } from "lodash";
import InfoLideresTemp from "../../components/InfoLideresTemp/InfoLideresTemp";
import CargarImagenes from "../../components/CargaImagenes/CargaImagenes";
import { getImagenes } from "../../api/lideres";

export default function LiderTemp(props) {
  const { openModal, setShowModal, lider } = props;
  const [signUpLoading, setSignUpLoading] = useState(false);
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
     

  const eliminarLiderTemp = () => {
    deleteLiderTemp(lider.id); 
    window.location.reload(); 
  }

    const onSubmit = e => {
        e.preventDefault();
        /* let validCount = 0;
        values(formData).some(value => {
            value && validCount++
            return null
        });

        if(validCount !== size(formData)) {
            toast.warning("Completa todo los campos del formulario");
        } else { */
            setSignUpLoading(true);
            signUpLider(lider).then(response => {
                if(response.code) {
                    toast.warning(response.message);
                } else {
                    toast.success("El registro ha sido correcto");
                }
            }).catch(() => {
                toast.error("Error del servidor, inténtelo más tarde!");
            }).finally(() => {
                setSignUpLoading(false);
            })                
    };

  /* const modificoUsuario = () => {
    modificarUsuario(user.id); 
    window.location.reload(); 
  } */
    return (
      <Row className="lista-lideresTemp__user">
        <Col id="usuario" sm={6}>
            <h2>{lider.nombre}</h2>
            <h3>{lider.departamento} - {lider.municipio} - {lider.territorio}</h3>
        </Col>
        <Col id="boton" sm={2}>
          <Button variant="info" 
            onClick={() => openModal(<
              InfoLideresTemp imagenes={imagenes} lider={lider} setShowModal={setShowModal}/>)
            }>
            Ver información
          </Button>
        </Col>
        <Col id="boton" sm={2}>
          <Button variant="warning" 
            onClick={() => openModal(<
              CargarImagenes lider={lider} setShowModal={setShowModal}/>)
            }>
            Cargar imagenes
          </Button>
        </Col>
        <Col id="boton" sm={2}>
          <Form onSubmit={onSubmit}>
            <Button variant="success" type="submit" onClick={() => eliminarLiderTemp()}>
              {!signUpLoading ? "Admitir" : <Spinner animation="border" />}
            </Button>
          </Form>          
        </Col>
      </Row>
    );
}