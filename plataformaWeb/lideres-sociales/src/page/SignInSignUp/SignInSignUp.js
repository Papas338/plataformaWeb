import React, { useState } from "react";
import { Container, Row, Col, Button } from "react-bootstrap";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faSearch,
  faComment,
  faLock,
} from "@fortawesome/free-solid-svg-icons";
import BasicModal from "../../components/Modal/BasicModal";
import SignUpForm from "../../components/SignUpForm/SignUpForm";
import SignInForm from "../../components/SignInForm/SignInForm";
import LogoWhite from "../../assests/png/paz.png";
import LogoTwittor from "../../assests/png/logo.png";
import Union from "../../assests/jpg/union.jpg";
import Ipazud from "../../assests/jpg/ipazud.jpg";
import Proximax from "../../assests/png/proximax.png";
import Logo from "../../assests/png/logo_40a_orgullo.png"
import "./SignInSignUp.scss";

export default function SignInSignUp(props) {
  const { setRefreshCheckLogin } = props;
  const [showModal, setShowModal] = useState(false);
  const [contentModal, setContentModal] = useState(null);
  const openModal = (content) => {
    setShowModal(true);
    setContentModal(content);
  };
  return (
    <>
      <Container className="signin-signup" fluid>
        <Row>
          <LeftComponent />
          <RightComponent
            openModal={openModal}
            setShowModal={setShowModal}
            setRefreshCheckLogin={setRefreshCheckLogin}
          />
        </Row>
      </Container>
      <BasicModal show={showModal} setShow={setShowModal}>
        {contentModal}
      </BasicModal>
    </>
  );
}

function LeftComponent() {
  return (
    <>
    <Col className="signin-signup__left" xs={6}>
      <div>
        <img src={Union} alt="Union" />
        <img src={LogoTwittor} alt="LogoTwittor" />
        <img src={Ipazud} alt="Ipazud" />
        <img src={Logo} alt="Logo" />
        <h5>
          <FontAwesomeIcon icon={faSearch} />
          Busca, consulta y descarga la información de los líderes sociales para recuperar y humanizar su memoria histórica.
        </h5>
        <h5>
          <FontAwesomeIcon icon={faLock} />
          La información de los líderes se encuentra respaldada gracias a la tecnología blockchain.
        </h5>
        <h5>
          <FontAwesomeIcon icon={faComment} />
          Habla y comparte la memoria de nuestros indígenas, no dejes que su legado se pierda en el tiempo.
        </h5>
      </div>      
    </Col>
    <Row>
      <Col id="proximax">
        <img src={Proximax} alt="Proximax" />  
      </Col>
    </Row>
    </>
  );
}

function RightComponent(props) {
  const { openModal, setShowModal, setRefreshCheckLogin } = props;
  return (
    <Col className="signin-signup__right" xs={6}>
      <div>
        <img src={LogoWhite} alt="LogoWhite" />
        <h2> Explora desde el mapa de Colombia todos los líderes que quieras consultar</h2>
        <Button
          variant="success"
          onClick={() => openModal(<SignUpForm setShowModal={setShowModal} />)}
        >
          Regístrate
        </Button>
        <Button
          variant="outline-primary"
          onClick={() =>
            openModal(
              <SignInForm setRefreshCheckLogin={setRefreshCheckLogin} />
            )
          }
        >
          Iniciar Sesión
        </Button>
      </div>
    </Col>
  );
}
