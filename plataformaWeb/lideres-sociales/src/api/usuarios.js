import { API_HOST } from "../utils/constant";
import { getTokenApi } from "./auth";

export function getUsuarios(paramsUrl) {
  const url = `${API_HOST}/listaUsuarios?${paramsUrl}`;

  const params = {
    headers: {
      Authorization: `Bearer ${getTokenApi()}`,
    },
  };

  return fetch(url, params)
    .then((response) => {
      return response.json();
    })
    .then((result) => {
      return result;
    })
    .catch((err) => {
      return err;
    });
}

export function modificarUsuario(ID) {
  
  const url = `${API_HOST}/modificarUsuario?id=${ID}`;

  const params = {
    method: "PUT",
    headers: {
      Authorization: `Bearer ${getTokenApi()}`,
    },
  };

  return fetch(url, params)
    .then((response) => {
      return response.json();
    })
    .then((result) => {
      return result;
    })
    .catch((err) => {
      return err;
    });
}
