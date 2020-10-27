import React from 'react';
import SignUpFormLider from "../../components/SignUpFormLider/SignUpFormLider";

import "./RegistroLideres.scss";

export default function RegistroLideres() {
    return (
        <>
            <div className="registro-lideres__header">
                <h1>Registro lideres</h1> 
            </div>           
            <SignUpFormLider />
        </>
    )
}
