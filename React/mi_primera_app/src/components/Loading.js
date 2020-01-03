import React from "react";

export default function Loading(props) {
    const {text} = props;
    return ( 
        <div>
            <h2>{text}</h2>
        </div>
    );
}