import React from "react";
import classes from "../css/chat.module.css";

class Message extends React.Component {
	constructor(props) {
		super(props);
		this.time = {
			hours: props.time.getHours() % 12,
			minutes: props.time.getMinutes() > 9 ? props.time.getMinutes() : "0" + props.time.getMintes(),
			seconds: props.time.getSeconds() > 9 ? props.time.getSeconds() : "0" + props.time.getSeconds()
		};
		this.message = {
			username: props.username,
			message: props.message,
		};
	}
	render() {
		return (
			<span className={classes.message}>
				<span className={classes.username}>
					{this.props.username}
				</span>
				<span className={classes.time}>
					{this.time.hours}:{this.time.minutes}:{this.time.seconds}
				</span>
				<span className={classes.userMessage}>{this.props.message}</span>
			</span>
		);
	}
}

export default Message;
