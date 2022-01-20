import React, { useState } from "react";

class SearchBar extends React.Component {
	constructor(props) {
		super(props);
	}

	render() {
		return (
			<div>
				<input
					type="text"
					placeholder="Search for class..."
					onChange={() => this.props.handleSearchMod}
					value={this.props.value}
				/>
			</div>
		);
	}
}

export default SearchBar;
