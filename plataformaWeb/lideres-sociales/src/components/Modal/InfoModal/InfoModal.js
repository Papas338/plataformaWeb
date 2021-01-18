import React from "react";
import { Modal } from "react-bootstrap";
import "./InfoModal.scss";

export default function InfoModal(props) {
  const { show, setShow, children } = props;

  return (
    <Modal
      className="infoModal"
      show={show}
      onHide={() => setShow(false)}
      centered
      size="lg"
    >
      <Modal.Body>{children}</Modal.Body>
    </Modal>
  );
}