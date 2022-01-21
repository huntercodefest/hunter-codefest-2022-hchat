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
			username: "daniel",
			askForNewUsername: true,
			room: new Room(null, null, 0, "Global"), // Temp
		};
	}
	handleUsernameChange = (event) => {
		event.preventDefault();
		this.setState({
			username: event.target[1].value,
			askForNewUsername: false,
			room: this.state.room,
		});
	}
	handleRoomChange = (Room) => {
		this.setState({
			...this.state,
			room: Room,
		});
	}
	render() {
		return (
			<div>
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
				/>
			</div>
		);
	}
}

export default App;
