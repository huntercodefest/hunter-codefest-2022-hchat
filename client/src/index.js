import React from "react";
import ReactDOM from "react-dom";
import UsernameInput from "./components/UsernameInput.js"
import classes from "./app.module.css";

import { useState } from "react";
// Get username as first action
// let username;
// do {
//     username = prompt();
// } while (!validateUsername(username));
let username = "test";
const room_num = "1";
let msg_format;
class App extends React.Component {
	constructor() {
		super();
		console.log("test");
		msg_format = `#${room_num}_${username}_`;
		console.log("Attempting WS connection");
		const ws = new WebSocket("ws://hchat.hopto.org:8080/ws");
		ws.onopen = () => {
			console.log("Connection made");
			ws.send(msg_format);
			ws.send(`${msg_format}${"hello server"}`);
		};
		ws.onmessage = (msg) => {
			console.log(`Server said ${msg}`);
		};
		ws.onerror = (error) => {
			console.log(`Socket error: ${error}`);
		};
	}
	render() {
		return (
			<div>
				<UsernameInput />

				<Msgform />
			</div>
		);
	}
}

function Msgform() {
	const [value, setValue] = useState(""); // value is equal to whatever the user inputs; see handleChange()

	const handleSubmit = (event) => {
		// Function called when user submits msg
		event.preventDefault(); // Stops browser from refreshing, which is the default for these events
		console.log(value);

		const text = value.trim(); // Removes whitespaces

		ChatBlock("UserTest", text); // Creates a ChatBlock

		if (text.length > 0) {
			setValue("");
		} // Clears the input value
	};

	const handleChange = (event) => {
		// Function called whenever <input> value is changed
		setValue(event.target.value);
	};

	// Message form JSX
	return (
		<div className={classes.chatBlock}>
			<form className="msgform" onSubmit={handleSubmit}>
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

// function Modal(props) {

//     function cancelHandler() {
//         props.onCancel();
//     }

//     function confirmHandler() {
//         props.onConfirm();
//     }

//     return (
//         <div className='modal'>
//             <p>Are you sure?</p>
//             <button className='btn btn--alt' onClick={cancelHandler}>Cancel</button>
//             <button className='btn' onClick={confirmHandler}>Confirm</button>
//         </div>
//     );
// }

// function Todo(props) {
//     const [ modalIsOpen, setModalIsOpen] = useState(false);

//     function deleteHandler() {
//         setModalIsOpen(true);
//     }
//     function closeModalHandler() {
//         setModalIsOpen(false);
//     }

//     return (
//         <div className='card'>
//             <h2>{props.text}</h2>
//             <div className='actions'>
//                 <button className='btn' onClick={deleteHandler}>Delete</button>
//             </div>
//             { modalIsOpen ? <Modal /> : null}
//         </div>
//     );
// }

// const sock = new WebSocket("ws://localhost:8080/ws")

function name() {}

function ChatBlock(props) {
	const { userName, messages } = props;

	const renderMessages = () => {
		console.log(userName + ":", messages);

		return (
			<div className="messageBlock" style={{ width: "100%" }}>
				<p>{userName}: {messages}</p>
			</div>
		);
	};

	return {renderMessages};
}

ReactDOM.render(<App />, document.getElementById("root"));
