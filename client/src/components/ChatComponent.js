import React from "react";
import Message from "./Message";
import ChatRoom from "./ChatRoom";
import MessageInput from "./MessageInput";
class ChatComponent extends React.Component {
	constructor(props) {
		// connect to server here
		super(props);
		this.ws = null;
		this.state = {
			messages: [],
			inputvalue: "",
		};
		// function binds
		this.handleChange = this.handleChange.bind(this);
		this.handleSubmit = this.handleSubmit.bind(this);
		this.sendToServer = this.sendToServer.bind(this);
		this.receiveFromServer = this.receiveFromServer.bind(this);
	}
	connectToServer() {
		// TODO
		this.ws.onopen = () => {
			this.ws.send(
				`#${this.props.room.room_num}_${this.props.username}:`
			);
		};
		this.ws.onmessage = (msg) => {
			this.receiveFromServer(msg);
		};
	}
	sendToServer(msg) {
		if (this.validateMessage(msg)) {
			this.ws.send(this.processMessage(msg));
		}
	}
	receiveFromServer(rcv_msg) {
		// TODO
		// Parse received message
		// And make new message component
		// with time and username and message
		const data = rcv_msg.data;
		console.log(this);
		let user = null;
		let msg = null;
		for (let i = 0; i < data.length; i++) {
			const user_end_index = data.indexOf(":");
			user = data.substring(data.indexOf("_") + 1, user_end_index);
			msg = data.substring(user_end_index + 1);
		}
		console.log(`received message ${msg} from server`);
		const time = new Date();
		const message = <Message time={time} username={user} message={msg} />;
		this.setState({
			...this.state,
			messages: this.state.messages.concat(message),
		});
	}
	validateMessage(msg) {
		// placeholder
		return msg.length > 0 && msg.length <= 500;
	}
	processMessage(msg) {
		return `#${this.props.room.room_num}_${this.props.username}:${msg}`;
	}
	handleSubmit(event) {
		event.preventDefault();
		this.sendToServer(event.target[0].value);
		this.setState({
			...this.state,
			inputvalue: "",
		});
		// Send message to server
	}
	handleChange(event) {
		this.setState({
			...this.state,
			inputvalue: event.target.value,
		});
	}
	attemptConnection() {
		// If ws connection already exists or username is empty
		if (this.ws !== null || this.props.username === "") {
			return;
		}
		this.ws = new WebSocket("ws://hchat.hopto.org:8080/ws");
		this.connectToServer();
	}
	render() {
		this.attemptConnection();
		return (
			<div>
				<ChatRoom
					room_desc={this.props.room.room_desc}
					messages={this.state.messages}
				/>
				<MessageInput
					handleSubmit={this.handleSubmit}
					handleChange={this.handleChange}
					value={this.state.inputvalue}
				/>
			</div>
		);
	}
}

export default ChatComponent;
