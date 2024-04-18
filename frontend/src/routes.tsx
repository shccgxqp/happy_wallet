import { createBrowserRouter } from "react-router-dom";
import Home from "./pages/home";
import Login from "./pages/login";
import Error from "./pages/error";
import Layout from "../src/Layout";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [
      {
        path: "",
        element: <Login />,
      },
      {
        path: "/home",
        element: <Home />,
      },
    ],
    errorElement: <Error />,
  },
]);
