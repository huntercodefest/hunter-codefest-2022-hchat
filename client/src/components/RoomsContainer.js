import React from "react";
import classes from "../css/rooms.module.css"

class RoomsContainer extends React.Component {
	constructor(props) {
		super(props);
	}
	jsx = [];
	process() {
		this.jsx = this.props.displayList.map((room) => (
			<button
				className={classes.RoomButton}
				key={room.room_desc}
				onClick={() => {
					this.props.handleRoomChange(room);
					this.props.handleListChange(room);
				}}>
				{room.room_desc}
			</button>
		));
	}
	render() {
		this.process();
		return <div className={classes.RoomsContainer}>{this.jsx}</div>;
	}
}

export default RoomsContainer;
