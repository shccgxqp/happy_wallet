import { createBrowserRouter } from "react-router-dom";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Error from "./pages/Error";
import Layout from "../src/pages/Layout";

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
