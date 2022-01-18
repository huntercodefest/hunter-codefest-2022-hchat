import React from "react";
import ChatBlock from "./ChatBlock";
import classes from "../app.module.css";

class MsgForm extends React.Component {
	constructor() {
		super();
		this.state = {
			value: "",
		};
	}
	// value is equal to whatever the user inputs; see handleChange()

	handleSubmit = (event) => {
		// Function called when user submits msg
		event.preventDefault(); // Stops browser from refreshing, which is the default for these events
		const text = this.state.value;
		console.log(text);
		if (this.validateMessage(text)) {
			// Creates a ChatBlock for the inputed message
			ChatBlock("UserTest", text);
			// Clears the input value
			this.setState({
				value: "",
			});
		}
	};

	handleChange = (event) => {
		// Function called whenever <input> value is changed
		console.log(event.target.value);
		this.setState({
			value: event.target.value,
		});
	};

	validateMessage(text) {
		// placeholder
		return text.length > 0;
	}
	// Message form JSX
	render() {
		return (
			<div className={classes.chatBlock}>
				<form className="msgform" onSubmit={this.handleSubmit}>
					<input
						className={classes.msginput}
						placeholder="Send a message..."
						onChange={this.handleChange}
						value={this.state.value}
					/>
				</form>
			</div>
		);
	}
}

export default MsgForm;
