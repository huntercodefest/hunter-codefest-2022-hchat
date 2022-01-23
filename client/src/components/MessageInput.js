import React from "react";
import classes from "../css/chat.module.css";

class MessageInput extends React.Component {
	constructor(props) {
		super(props);
	}
	// value is equal to whatever the user inputs; see handleChange()

	// Message form JSX
	render() {
		return (
			<div className={classes.div}>
				<form
				className="msgform"
					onSubmit={(event) => this.props.handleSubmit(event)}
					id="sub_msg">
					<input
						className={classes.chatInput}
						placeholder="Send a message..."
						onChange={(event) => this.props.handleChange(event)}
						value={this.props.value}
					/>
				</form>
				<button form="sub_msg" className={classes.button}>
					Send
				</button>
			</div>
		);
	}
}

export default MessageInput;
