import Home from "../page/Home";
import AprobarUsuarios from "../page/AprobarUsuarios";
import Error404 from "../page/Error404";

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
    path: "*",
    page: Error404,
  },
];
