import React from "react";
import arrow_path from "../backarrow.svg";
import classes from "../css/rooms.module.css"
class BackButton extends React.Component {
	constructor(props) {
		super(props);
	}
	render() {
		return (
			<img
				className={`${classes.backButton} ${this.props.modifyclass}`}
				src={arrow_path}
				alt={"backbutton"}
				onClick={() => this.props.handleBackClick()}
			/>
		);
	}
}

export default BackButton;
