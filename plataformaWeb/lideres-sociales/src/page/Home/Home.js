import React from "react";
import { Container, Row, Col, Button } from "react-bootstrap";
import { Link } from "react-router-dom";
import Logo from "../../assests/png/logo_40a_orgullo.png";
import Banner from "../../assests/jpg/banner.jpg";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
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

  const { setRefreshCheckLogin } = props;
  console.log(props)

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
          <Actions />
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
        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi sodales
        elit eget nibh pellentesque, et porta sapien maximus. Proin interdum
        dolor eget velit volutpat accumsan. Aenean blandit facilisis mi eget
        congue. Integer ipsum orci, pellentesque id augue eget, faucibus cursus
        odio. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices
        posuere cubilia curae; Donec at velit blandit, semper ipsum quis,
        vestibulum mauris. Integer sed varius velit. Proin eu lorem tellus.
        Curabitur tincidunt, tellus eu tempus tincidunt, risus mi lacinia nulla,
        non fringilla purus neque at metus. Praesent semper pretium ligula, sit
        amet mollis velit imperdiet ut. Suspendisse sagittis sapien eget sapien
        viverra, quis consequat neque viverra.
      </p>
    </Col>
  );
}

function Actions() {
  return (
    <Col className="home__actions" xs={6}>
      <div>
        <Button variant="success" as={Link}>
          Ver Líderes
        </Button>
        <Button variant="success" as={Link} to="/registroLideres">
          Registro Líderes
        </Button>
        <Button variant="success" as={Link} to="/aprobarUsuarios">
          Aprobar Usuarios (Pasantes)
        </Button>
        <Button variant="success" as={Link}>
          Listado de Líderes
        </Button>
      </div>
    </Col>
  );
}
