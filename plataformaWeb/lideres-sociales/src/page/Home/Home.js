import React from "react";
import { Container, Row, Col, Button } from "react-bootstrap";
import { Link } from "react-router-dom";
import Logo from "../../assests/png/logo_40a_orgullo.png";
import Banner from "../../assests/jpg/banner.jpg";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import userAuth from "../../hooks/userAuth";
import { logoutApi } from "../../api/auth";
import {
  faSoundcloud,
  faFacebook,
  faTwitter,
  faYoutube,
} from "@fortawesome/free-brands-svg-icons";

import "./Home.scss";

library.add(faFacebook, faTwitter, faYoutube, faSoundcloud);

export default function Home(props) {

  const user = userAuth();

  var userRole = user.role;

  const { setRefreshCheckLogin } = props;

  const logout = () => {
    logoutApi();
    setRefreshCheckLogin(true);
  }

  return (
    <>
      <Container>
        <Row className="home__bar">
          <Col>
            <div id="redes">
              <FontAwesomeIcon icon={faFacebook} />
              <FontAwesomeIcon icon={faTwitter} />
              <FontAwesomeIcon icon={faYoutube} />
              <FontAwesomeIcon icon={faSoundcloud} />
            </div>
            <div id="cerrar">
              <Button as={Link} variant="danger" onClick={logout}>Cerrar Sesión</Button>
            </div>
          </Col>
        </Row>
        <Row>
          <HeaderLogo />
          <HeaderBanner />
        </Row>
        <Row>
          <Text />
          <Actions role={userRole} />
        </Row>
      </Container>
    </>
  );
}

function HeaderLogo() {
  return (
    <Col className="home__headerLogo" sm={4}>
      <div>
        <img src={Logo} alt="onic" />
      </div>
    </Col>
  );
}

function HeaderBanner() {
  return (
    <Col className="home__headerBanner" sm={8}>
      <div>
        <img src={Banner} alt="indigenas" />
      </div>
    </Col>
  );
}

function Text() {
  return (
    <Col className="home__text" xs={6}>
      <p>
        Los líderes indígenas realizan más actividades que no necesariamente están relacionadas con lo que compete a los asuntos indígenas. Los líderes también trabajan temas relacionados con construcción y mejoramiento de vías, cultivos, entre otros. Es importante entender que un líder trabaja desde diferentes frentes, sin importar que se catalogue desde uno solo. 
      </p>
      <p id="parrafo">
        Es importante la labor que ejerce un líder social indígena sobre su territorio, sobre recuperar la tierra que han perdido, denunciar actos delictivos, además de conservar la memoria histórica de sus ancestros, y la muerte de estos genera un impacto directo en los derechos humanos y en la protección de los mismos, en las causas que protegían, en su forma de transmitir la cultura, en las colectivos, organizaciones y comunidades que representaban, aumentando así la percepción de vulnerabilidad, y de indefensión de los derechos humanos.
      </p>
    </Col>
  );
}

function Actions(userRole) {
  
  if (userRole.role == "Administrador") {
    return (
      <Col className="home__actions" xs={6}>
        <div>
          <Button variant="success" as={Link} to="/visualizarLideres">
            Ver Líderes
          </Button>
          <Button variant="success" as={Link} to="/informacionLideres">
            Listado de Líderes
          </Button>
          <Button variant="success" as={Link} to="/registroLideres">
            Registro Líderes
          </Button>
          <Button variant="success" as={Link} to="/verificacionLideres">
            Verificar Líderes
          </Button>
          <Button variant="success" as={Link} to="/aprobarUsuarios">
            Aprobar Usuarios (Pasantes)
          </Button>        
        </div>
      </Col>
    );
  } else if (userRole.role == "Pasante") {
    return (
      <Col className="home__actions" xs={6}>
        <div>
          <Button variant="success" as={Link}>
            Ver Líderes
          </Button>
          <Button variant="success" as={Link} to="/informacionLideres">
            Listado de Líderes
          </Button>
          <Button variant="success" as={Link} to="/registroLideres">
            Registro Líderes
          </Button>        
        </div>
      </Col>
    );
  } 

  return (
    <Col className="home__actions" xs={6}>
      <div>
        <Button variant="success" as={Link}>
          Ver Líderes
        </Button>
        <Button variant="success" as={Link} to="/informacionLideres">
          Listado de Líderes
        </Button>      
      </div>
    </Col>
  );
}
