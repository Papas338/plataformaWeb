import Home from "../page/Home";
import AprobarUsuarios from "../page/AprobarUsuarios";
import Error404 from "../page/Error404";
import RegistroLideres from "../page/RegistroLideres/RegistroLideres";

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
    path: "*",
    page: Error404,
  },
];
