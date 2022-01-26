import React from "react";
import { Schools } from "./Rooms"
import classes from "../css/rooms.module.css"

class RoomsContainer extends React.Component {
	constructor(props) {
		super(props);
		this.schoollist = Schools.CurrList.map(school => school)
	}
	jsx = [];
	process() {
		this.jsx = this.props.displayList.map((room) => {
			let styles = `${classes.roomButton}`;
			let room_desc = room.room_desc
			// If global room
			if (room.school === "_GLOB"){
				styles = `${styles} ${classes.displayorder} ${classes.bold}`;
			}
			// If school name is same as room name
			else if (this.schoollist.map(school => school.school).includes(room_desc)) {
				// Set description to the full school name
				room_desc = this.schoollist.filter(
					(school) => school.school === room_desc)[0].room_desc
				console.log(room_desc)
				// Add order: 1 property to styles
				styles = `${styles} ${classes.displayorder} ${classes.bold}`;
				console.log(styles)
			}
			else if (this.props.listDesc === "classes" && room.major === room.room_desc) {
				console.log(this.props.listDesc)
				console.log(true)
				console.log(room.major + " " + room.room_desc)
				styles = `${styles} ${classes.displayorder} ${classes.bold}`;
				console.log(styles)
			}
			return <button
				className={`${styles}`}
				key={room.room_desc}
				onClick={() => {
					this.props.handleRoomChange(room);
					this.props.handleListChange(room);
				}}>
				{room_desc}
			</button>
		});
	}
	render() {
		this.process();
		return <div className={classes.roomsContainer}>{this.jsx}</div>;
	}
}

export default RoomsContainer;
