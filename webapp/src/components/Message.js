import React from "react";
import classes from "../css/chat.module.css";

class Message extends React.Component {
	constructor(props) {
		super(props);
		this.time = {
			hours: props.time.getHours() % 12,
			minutes: props.time.getMinutes(),
			seconds: props.time.getSeconds(),
		};
		this.message = {
			username: props.username,
			message: props.message,
		};
	}
	render() {
		return (
			<span classesName={classes.message}>
				<span className={classes.username}>
					{this.props.username}
				</span>
				<span className={classes.time}>
					{this.time.hours}:{this.time.minutes}:{this.time.seconds}
				</span>
				<br/>
				<span className={classes.userMessage}>{this.props.message}</span>
				<br/>
			</span>
		);
	}
}

export default Message;
