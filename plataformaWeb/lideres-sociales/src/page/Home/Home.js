import React from "react";
import { Container, Row, Col, ButtonGroup, Button } from "react-bootstrap";
import { Link } from "react-router-dom";
import Logo from "../../assests/png/logo_40a_orgullo.png";
import Banner from "../../assests/jpg/banner.jpg";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faGithubAlt,
  faGoogle,
  faFacebook,
  faTwitter,
  faYoutube,
} from "@fortawesome/free-brands-svg-icons";

import "./Home.scss";

library.add(faGithubAlt, faGoogle, faFacebook, faTwitter, faYoutube);

export default function Home() {
  return (
    <>
      <Row>
        <Bar />
      </Row>
      <Container>
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

function Bar() {
  return (
    <Col className="home__bar">
      <div>
        <FontAwesomeIcon icon={faFacebook} />
        <FontAwesomeIcon icon={faTwitter} />
        <FontAwesomeIcon icon={faYoutube} />
      </div>
    </Col>
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
        <Link>Ver Líderes</Link>
        <Link>Cargar Líderes</Link>
        <Link to="/aprobarUsuarios">Aprobar Usuarios (Pasantes)</Link>
        <Link>Listado de Líderes</Link>
      </div>
    </Col>
  );
}
