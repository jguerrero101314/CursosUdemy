import React from "react";

export default function Buttons(props){
    const {fnShowName} = props;
    console.log(props);

    return(
        <div>
            <button onClick={()=> fnShowName("Joel Guerrero")}>Click</button>
        </div>


    );

}