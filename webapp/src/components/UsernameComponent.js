import React from "react";
import Modal from "./Modal"
import classes from "../css/app.module.css";

class UsernameComponent extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			value: "",
		};
		this.handleValueChange = this.handleValueChange.bind(this)
	}
	handleValueChange = (event) => {
		this.setState({
			// target is form child 1 is input
			value: event.target.value,
		});
	}
	render() {
		return (
			<Modal
				handleUsernameChange={this.props.handleUsernameChange}
				handleValueChange={this.handleValueChange}
				value={this.state.value}
			/>
		);
	}
}

export default UsernameComponent;

/* render() {
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
 */
