import React from "react";
//import Loading from "./components/Loading";
import ListUsers from "./components/ListUsers";
import ListUsersV2 from "./components/ListUsersV2";
import Buttons from "./components/Buttons";
import {ReactComponent as ReactLogo} from "./assets/react.svg";


export default  function App() {
    // const textObjet = {
    //     title: "Cargando",
    //     time: 6
    // };

    // const users = [
    //     "Joel Guerrero",
    //      "Ruth Andino",
    //      "Nico Andino"
    //     ];
    // const employees = [
    //     {name : "Joel Guerrero", years: 22},
    //     {name : "Carlos Perez", years: 29},
    //     {name : "Raul Ruiz", years: 16},
    //     {name : "Ruth Andino", years: 17}
    // ]
    const showNameConsole = (name) =>{
        console.log(name);
    }
    return ( 
    //    <div className="App">
    //        <h1>Mi App</h1>
    //        <Loading text ="Cargando todo..."/>
    //    </div>

    //    <div className="App">
    //        <h1>Mi App</h1>
    //        <Loading data = {textObjet}/>
    //    </div>

            // <div className="App">
            //     <h1>Mi App</h1>
            //     <Buttons fnShowName = {showNameConsole}/>
            //     {/* <ListUsers  ListUsers = {users}/>
            //     <ListUsersV2  users = {employees}/> */}
            // </div>

                  <div className="App">
                        <h1>Mi App</h1>
                        <ReactLogo style={{width:80, height:80}}/>
                 </div>

    
    );
}

