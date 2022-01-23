import React from "react";
import UsernameComponent from "./UsernameComponent";
import RoomsComponent from "./RoomsComponent";
import ChatComponent from "./ChatComponent";
import classes from "../css/app.module.css";
import { Schools } from "./Rooms";
import { Room } from "./Rooms";

class App extends React.Component {
	constructor() {
		super();
		this.state = {
			username: "",
			askForNewUsername: true,
			clear: false,
			room: Schools.CurrList[0], // Global room
		};
		this.handleUsernameChange = this.handleUsernameChange.bind(this);
		this.handleRoomChange = this.handleRoomChange.bind(this);
		this.clearedChatRoom = this.clearedChatRoom.bind(this);
	}
	handleUsernameChange = (event) => {
		event.preventDefault();
		const new_username = event.target[0].value;
		if (this.validateUsername(new_username)) {
			this.setState({
				username: event.target[0].value,
				askForNewUsername: false,
				room: this.state.room,
			});
		} else {
			alert(
				"Username must be between 3 and 20 characters long\nUsername can only contain letters, numbers, or underscores"
			);
		}
	};
	handleRoomChange = (room) => {
		if (room.room_num) {
			console.log(`Changed to room#${room.room_num}`);
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
			clear: false,
		});
	};

	validateUsername = (username) => {
		// test input length and character matching (only letters, numbers, or dashes)
		const regex = new RegExp("^[A-Za-z0-9_-]{3,20}$");
		return regex.test(username);
	};
	render() {
		return (
			<div className={classes.app}>
				{/* {this.state.askForNewUsername ? (
					<UsernameComponent
						handleUsernameChange={this.handleUsernameChange}
						room={this.state.room}
					/>
				) : null} */}
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
