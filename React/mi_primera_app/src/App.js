import React from "react";
import Loading from "./components/Loading";


export default  function App() {
    return ( 
       <div className="App">
           <h1>Mi App</h1>
           <Loading text ="Cargando todo..."/>
       </div>
    );
}

