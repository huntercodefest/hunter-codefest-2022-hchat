import React from "react";
import UsernameComponent from "./UsernameComponent";
import RoomsComponent from "./RoomsComponent";
import ChatComponent from "./ChatComponent";
import classes from "../css/app.module.css";
import { Room } from "./Rooms";

class App extends React.Component {
	constructor() {
		super();
		this.state = {
			username: "",
			askForNewUsername: true,
			clear: false,
			room: new Room(null, null, 0, "Global"), // Temp
		};
		this.handleUsernameChange = this.handleUsernameChange.bind(this);
		this.handleRoomChange = this.handleRoomChange.bind(this);
		this.clearedChatRoom = this.clearedChatRoom.bind(this)
	}
	handleUsernameChange = (event) => {
		event.preventDefault();
		this.setState({
			username: event.target[0].value,
			askForNewUsername: false,
			room: this.state.room,
		});
	};
	handleRoomChange = (room) => {
		if (room.room_num) {
			console.log(`Changed to room#${room.room_num}`)
			this.setState({
				...this.state,
				clear: true,
				room: room,
			});
		}
	};

	clearedChatRoom = () => {
		this.setState({
			...this.state,
			clear: false
		})
	}
	render() {
		return (
			<div>
				{this.state.askForNewUsername ? (
					<UsernameComponent
						handleUsernameChange={this.handleUsernameChange}
						room={this.state.room}
					/>
				) : null}
				<RoomsComponent handleRoomChange={this.handleRoomChange} />
				<ChatComponent
					username={this.state.username}
					room={this.state.room}
					clear={this.state.clear}
					clearFunc={this.clearedChatRoom}
				/>
			</div>
		);
	}
}

export default App;
