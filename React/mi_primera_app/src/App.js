import React from "react";
//import Loading from "./components/Loading";
import ListUsers from "./components/ListUsers";
import ListUsersV2 from "./components/ListUsersV2";


export default  function App() {
    // const textObjet = {
    //     title: "Cargando",
    //     time: 6
    // };

    const users = [
        "Joel Guerrero",
         "Ruth Andino",
         "Nico Andino"
        ];
    const employees = [
        {name : "Joel Guerrero", years: 22},
        {name : "Carlos Perez", years: 29},
        {name : "Raul Ruiz", years: 16},
        {name : "Ruth Andino", years: 17}
    ]
    return ( 
    //    <div className="App">
    //        <h1>Mi App</h1>
    //        <Loading text ="Cargando todo..."/>
    //    </div>

    //    <div className="App">
    //        <h1>Mi App</h1>
    //        <Loading data = {textObjet}/>
    //    </div>

            <div className="App">
                <h1>Mi App</h1>
                <ListUsers  ListUsers = {users}/>
                <ListUsersV2  users = {employees}/>
            </div>

    
    );
}

