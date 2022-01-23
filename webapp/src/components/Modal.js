import React from "react";
import classes from "../css/modal.module.css";
class Modal extends React.Component {
	constructor(props) {
		super(props);
	}
	render() {
		return (
			<div className={classes.backdrop}>
				<div className={classes.modal}>
					<form
						onSubmit={(event) => {
							this.props.handleUsernameChange(event);
						}}>
						<h2>Enter Username</h2>
						<input
							type="text"
							onChange={(event) =>
								this.props.handleValueChange(event)
							}
						/>
					</form>
				</div>
			</div>
		);
	}
}

export default Modal;
