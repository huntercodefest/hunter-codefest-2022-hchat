import React from "react";
import { Room, RoomList, Schools } from "./Rooms";

class RoomsComponent extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			searchList: Schools,
			searchValue: "",
			displayList: this.state.searchList.CurrList,
			room: new Room("", "", 0, "Global"),
		};
	}
	handleSearchMod(event) {
		const searchValue = event.target.value;

		// modify displaylist
		// change value
		this.setState({
			...this.state,
			searchValue: event.target.value,
			displayList: this.state.searchList.CurrList.filter((room) =>
				room.room_desc.contains(searchValue)
			),
		});
	}
	handleBackClick() {
		// modify searchlist
	}
	render() {
		return (
			<div>
				<BackButton handleBackClick={this.handleBackClick} />
				<SearchBar
					searchList={this.state.searchList.CurrList}
					value={this.state.searchValue}
					handleSearchMod={this.handleSearchMod}
				/>
				<p>Found {this.displayList.length} rooms</p>
				<Room_List
					displayList={this.state.displayList}
					handleRoomChange={this.props.handleRoomChange}
				/>
			</div>
		);
	}
}

export default RoomsComponent;
