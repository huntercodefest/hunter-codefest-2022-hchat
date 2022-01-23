import React from "react";
import classes from "../css/chat.module.css";
class ChatRoom extends React.Component {
	constructor(props) {
		super(props);
	}
	render() {
		return (
			<div className={classes.chatModule}>
				<p>Chatroom starts here</p>
				<span>{this.props.room_desc}</span>
				<div>{this.props.messages}</div>
			</div>
		);
	}
}

export default ChatRoom;
