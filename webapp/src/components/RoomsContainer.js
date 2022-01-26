import React from "react";
import classes from "../css/rooms.module.css"

class RoomsContainer extends React.Component {
	constructor(props) {
		super(props);
	}
	jsx = [];
	process() {
		this.jsx = this.props.displayList.map((room) => {
			const bold = room.room_num ? classes.bold : null;
			return <button
				className={`${classes.roomButton} ${bold}`}
				key={room.room_desc}
				onClick={() => {
					this.props.handleRoomChange(room);
					this.props.handleListChange(room);
				}}>
				{room.room_desc}
			</button>
			});
	}
	render() {
		this.process();
		return <div className={classes.roomsContainer}>{this.jsx}</div>;
	}
}

export default RoomsContainer;
