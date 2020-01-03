import React from "react";

export default function ListUsers(props){
    const {ListUsers} = props;

    return(
        <div>
            <h4>Lista de usuarios</h4>
            <ul>
               {ListUsers.map((name, index) =>(
                   <li key={index}>{name}</li>

               ))} 
            </ul>
        </div>

    );
}