import React from "react";

class Room_List extends React.Component {
	constructor(props) {
		super(props);
	}
	jsx = [];
	process() {
		this.jsx = this.props.displayList.map((room) => (
			<button
				key={room.room_num}
				onClick={() => this.props.handleRoomChange(room)}>
				{room.room_desc}
			</button>
		));
	}
	render() {
		this.process();
		return <div>{this.jsx}</div>;
	}
}

export default Room_List