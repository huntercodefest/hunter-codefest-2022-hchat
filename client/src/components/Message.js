import React from "react";

class Message extends React.Component {
	constructor(props) {
		super(props);
		this.time = {
			hours: props.time.getHours(),
			minutes: props.time.getMinutes(),
		};
        this.message = {
            username: props.username,
            message: props.message,
        }
	}
	render() {
		return (
        <p>{this.time.hours}:{this.time.minutes}</p>
        );
	}
}

export default Message;
