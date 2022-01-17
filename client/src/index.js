import React from "react";
import ReactDOM from "react-dom";
import classes from "./test.module.css";

import { useState } from 'react';
// Get username as first action
// let username;
// do {
//     username = prompt();
// } while (!validateUsername(username));
let username = "test";
const room_num = "1"
let msg_format;
class App extends React.Component {
    constructor() {
        super();
        console.log("test")
        msg_format = `#${room_num}_${username}_`;
        console.log("Attempting WS connection")
        const ws = new WebSocket("ws://hchat.hopto.org:8080/ws")
        ws.onopen = () => {
            console.log("Connection made")
            ws.send(msg_format)
            ws.send(`${msg_format}${"hello server"}`)
        }
        ws.onmessage = (msg) =>{
            console.log(`Server said ${msg}`)
        }
        ws.onerror = (error) =>{
            console.log(`Socket error: ${error}`)
        }
    }
    render() {
        return (
            <div>
                <Msgform />
            </div>
        )
    }
}

function Msgform() {
    const [value, setValue] = useState('');   // value is equal to whatever the user inputs; see handleChange()

    const handleSubmit = (event) => {
        // Function called when user submits msg
        event.preventDefault();  // Stops react from re-rendering (default)
        console.log(value);

        const text = value.trim()  // Removes whitespaces

        if (text.length > 0) {setValue('')}  // Clears the input value
    }

    const handleChange = (event) => {
        // Function called whenever <input> value is changed
        setValue(event.target.value);
    }

    // Message form JSX 
    return (
        <div className={classes.chatblock}>
            <form className='msgform' onSubmit={handleSubmit}>
                <input 
                    className={classes.msginput}
                    placeholder="Send a message..."
                    onChange={handleChange}
                    onSubmit={handleSubmit}
                    value={value}
                />
            </form>
        </div>
    );
}

// wait for room
/*
...
*/
// attempt connection to server
function validateUsername(username){
    // placehilder
    return username !== ""
}

// const sock = new WebSocket("ws://localhost:8080/ws")
ReactDOM.render(<App />, document.getElementById("root"));



