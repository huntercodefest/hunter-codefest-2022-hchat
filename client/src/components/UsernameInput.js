import React from "react";
import classes from "../app.module.css";

class UsernameInput extends React.Component {
	constructor() {
		super();
		this.state = {
			vis: true,
			inputValue: "",
		};
	}

	render() {
		return (
			<div
				className={
					this.state.vis ? classes.backdrop : classes.displaynone
				}>
				<div className={classes.modal}>
					<h2>Enter Username</h2>
					<form onSubmit={(event) => this.handleSubmit(event)}>
						<input
							type="text"
							placeholder="Enter Username..."
							value={this.state.inputValue}
							onChange={(event) => this.handleChange(event)}
						/>
					</form>
				</div>
			</div>
		);
	}
	validateUsername(username){
		// placeholder
		return username !== ""
	}
	handleSubmit(event) {
		event.preventDefault();
		const val = event.target.firstElementChild.value;
		if (this.validateUsername(val)) {
			this.setState({
				vis: false,
				inputValue: "",
			});
		}
	}
	handleChange(event) {
		console.log(event);
		const val = event.target.value;
		this.setState({
			vis: true,
			inputValue: val,
		});
	}
}

export default UsernameInput;                              