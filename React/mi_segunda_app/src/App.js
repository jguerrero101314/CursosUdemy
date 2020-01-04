import React from "react";
import { Button, Accordion, Card, Alert } from "react-bootstrap";
import "./App.css";

export default function App() {
  return (
    <div className="App">
      <h1>React Boostrap</h1>
      <Button variant="primary" onClick={() => console.log("Hola")}>
        Primary
      </Button>
      <Alert variant="danger">Peligro!!!</Alert>
    </div>
  );
}
