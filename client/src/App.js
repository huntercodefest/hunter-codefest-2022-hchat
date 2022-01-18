import React from "react"
import UsernameInput from "./components/UsernameInput"
import MsgForm from "./components/MsgForm"
import classes from "./app.module.css";

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
                {/* Get username as first action */}
				<UsernameInput />

				<MsgForm />
			</div>
		);
	}
}

export default App