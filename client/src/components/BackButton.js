import React from "react";
import arrow_path from "../backarrow.svg";

class BackButton extends React.Component {
	constructor(props) {
		super(props);
	}
	render() {
		return (
			<img
				src={arrow_path}
				onClick={() => this.props.handleBackClick()}
			/>
		);
	}
}

export default BackButton;
