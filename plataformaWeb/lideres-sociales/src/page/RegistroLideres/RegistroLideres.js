import React from 'react';
import SignUpFormLider from "../../components/SignUpFormLider/SignUpFormLider";
import userAuth from "../../hooks/userAuth";

import "./RegistroLideres.scss";

export default function RegistroLideres() {
    const user = userAuth();
    if (user.role == "Administrador" || user.role == "Pasante") {
        return (
            <>
                <div className="registro-lideres__header">
                    <h1>Registro lideres</h1> 
                </div>           
                <SignUpFormLider />
            </>
        )    
    } else {
        return (
            <>
                <div className="registro-lideres__header">
                    <h1>Acceso denegado</h1> 
                </div>
            </>
        )        
    }    
}
