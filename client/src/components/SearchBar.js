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
					onChange={(event) => this.props.handleSearchMod(event)}
					value={this.props.value}
				/>
			</div>
		);
	}
}

export default SearchBar;
