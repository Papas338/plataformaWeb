import { API_HOST } from "../utils/constant";
import { getTokenApi } from "./auth";

export function getDepartamentos(nombre) {
    const url = `${API_HOST}/departamentos?search=${nombre}`;
  
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

  export function getMunicipios(departamento, municipio) {
    const url = `${API_HOST}/municipios?departamento=${departamento}&municipio=${municipio}`;
  
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