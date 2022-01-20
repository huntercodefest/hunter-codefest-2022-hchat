import React from "react";

class Message extends React.Component {
	constructor(props) {
		super(props);
		this.time = {
			hours: props.time.getHours() % 12,
			minutes: props.time.getMinutes(),
		};
		this.message = {
			username: props.username,
			message: props.message,
		};
	}
	render() {
		return (
			<span>
				<p>{this.props.username}</p>
				<p>{this.props.message}</p>
				<p>
					{this.time.hours}:{this.time.minutes}
				</p>
			</span>
		);
	}
}

export default Message;
