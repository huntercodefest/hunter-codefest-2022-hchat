import React from "react";

class RoomsContainer extends React.Component {
	constructor(props) {
		super(props);
	}
	jsx = [];
	process() {
		this.jsx = this.props.displayList.map((room) => (
			<button
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
		return <div>{this.jsx}</div>;
	}
}

export default RoomsContainer;
