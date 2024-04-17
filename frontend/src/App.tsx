import React from 'react';
import { RouterProvider } from 'react-router-dom';
import { router } from "./routes";
import Layout from "../src/Layout";
import './App.css';

function App() {
  return (
    <div className="App">
      <Layout>
              <RouterProvider router={router} />
      </Layout>
    </div>
  );
}

export default App;
