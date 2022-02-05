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
			<div className={classes.chatOuterUI}>
				<form
					onSubmit={(event) => this.props.handleSubmit(event)}
					id="sub_msg">
					<input
						className={classes.chatInput}
						placeholder={`Send a message in ${this.props.room_desc}...`}
						onChange={(event) => this.props.handleChange(event)}
						value={this.props.value}
					/>
				</form>
			</div>
		);
	}
}

export default MessageInput;
