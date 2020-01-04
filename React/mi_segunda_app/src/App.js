import React from "react";
//import { Button, Accordion, Card, Alert } from "react-bootstrap";
import Car from "./components/Car";
import "./App.css";

export default function App() {
  //   return (
  //     <div className="App">
  //       <h1>React Boostrap</h1>
  //       <Button variant="primary" onClick={() => console.log("Hola")}>
  //         Primary
  //       </Button>
  //       <Alert variant="danger">Peligro!!!</Alert>
  //     </div>
  //   );
  return(
      <div>
          <h1>React Hook useState</h1>
        
          <Car/>
      </div>
  );
}
