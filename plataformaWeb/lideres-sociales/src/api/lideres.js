import { API_HOST } from "../utils/constant";
import { getTokenApi } from "./auth";

export function signUpLider(lider) {
    console.log(lider);
    const url = `${API_HOST}/registroLider`;
    const params = {
      method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${getTokenApi()}`,
    },
    body: JSON.stringify(lider),    
    };

    console.log(params);

    return fetch(url, params).then(response => {
      if(response.status >= 200 && response.status < 300) {
        console.log(response.body)
        return response.json();
      }
      return {code: 404, message: "Datos de lider no validos"}
    }).then(result => {
      return result;
    }).catch(err => {
      return err;
    })
  }