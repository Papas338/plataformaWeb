import Home from "../page/Home";
import AprobarUsuarios from "../page/AprobarUsuarios";
import Error404 from "../page/Error404";
import RegistroLideres from "../page/RegistroLideres/RegistroLideres";
import InformacionLideres from "../page/InformacionLideres/InformacionLideres";
import VerificarLideres from "../page/VerificarLideres/VerificarLideres";
import VisualizarLideres from "../page/VisualizarLideres/VisualizarLideres";

export default [
  {
    path: "/",
    exact: true,
    page: Home,
  },
  {
    path: "/aprobarUsuarios",
    exact: true,
    page: AprobarUsuarios,
  },
  {
    path: "/registroLideres",
    exact: true,
    page: RegistroLideres,
  },
  {
    path: "/informacionLideres",
    exact: true,
    page: InformacionLideres,
  },
  {
    path: "/verificacionLideres",
    exact: true,
    page: VerificarLideres,
  },
  {
    path: "/visualizarLideres",
    exact: true,
    page: VisualizarLideres,
  },
  {
    path: "*",
    page: Error404,
  },
];
