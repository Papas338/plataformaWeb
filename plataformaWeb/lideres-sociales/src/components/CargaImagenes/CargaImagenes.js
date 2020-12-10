import React, { useState, useCallback } from 'react';
import { toast } from "react-toastify";
import { subirImagen, registroImagen } from "../../api/lideres";
import { Form, Col, Row, Button, Spinner, Container } from 'react-bootstrap';
import { useDropzone } from "react-dropzone";
import { Camera } from "../../utils/icons";

import "./CargaImagenes.scss";

export default function CargaImagenes(props) {
    const {lider, setShowModal} = props;
    const [imagen, setImagen] = useState(null);
    const [imageFile, setImageFile] = useState(null);
    const [signUpLoading, setSignUpLoading] = useState(false);
    
    const onDropImagen = useCallback(acceptedFile => {
        const file = acceptedFile[0];
        setImagen(URL.createObjectURL(file));
        setImageFile(file);
    })

    const { getRootProps, getInputProps } = useDropzone({
        accept: "image/jpeg, image/png",
        noKeyboard: true,
        multiple: false,
        onDrop: onDropImagen
    })

    const onSubmit = e => {
        e.preventDefault();

        if(imageFile) {
            subirImagen(imageFile)
            .then(() => {
                registroImagen(lider.id).catch(() => {
                    toast.error("Error al subir la imagen");
                });
            })
            .catch(() => {
                toast.error("Error al subir la imagen");
            });            
            toast.success("Imagen cargada");
            setSignUpLoading(true);
            setShowModal(false);
        }
        setSignUpLoading(false);
    }

    return (
        <>
            <Container>
                <Form onSubmit={onSubmit}>
                    <Row 
                        className="imagen" 
                        style={{ backgroundImage: `url('${imagen}')` }}
                        {...getRootProps()}
                    >
                        <input {...getInputProps()} />
                        <Camera />
                    </Row>
                    <Row>
                        <Col md="auto" className="boton_imagen">
                            <Button variant="success" type="submit">
                                {!signUpLoading ? "Registrar Lider" : <Spinner animation="border" />}
                            </Button>
                        </Col>
                    </Row>                
                </Form>
            </Container>
        </>
    )
}
