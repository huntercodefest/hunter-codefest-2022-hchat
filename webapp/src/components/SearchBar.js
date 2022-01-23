import React from "react";
import classes from "../css/rooms.module.css"
class SearchBar extends React.Component {
	constructor(props) {
		super(props);
	}

	render() {
		return (
			<div>
				<input
					type="text"
					placeholder="Search for a class..."
					className={classes.searchBar}
					onChange={(event) => this.props.handleSearchMod(event)}
					value={this.props.value}
				/>
			</div>
		);
	}
}

export default SearchBar;
