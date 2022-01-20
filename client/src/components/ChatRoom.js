import React from "react";

class ChatRoom extends React.Component {
	constructor(props) {
		super(props);
	}

	render() {
		return (
			<div>
				<span>{this.props.Room_Desc}</span>
				<div>{this.props.messages}</div>
			</div>
		);
	}
}

export default ChatRoom;
