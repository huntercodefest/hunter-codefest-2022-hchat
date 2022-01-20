import React from "react";
import UsernameComponent from "./UsernameComponent";
import MessageInput from "./MessageInput";
import classes from "./app.module.css";
import { Room } from "./Rooms";

class App extends React.Component {
	constructor() {
		super();
		this.state = {
			username: "",
			askForNewUsername: true,
			Room: new Room(null, null, 0, "Global"), // Temp
		};
	}
	handleUsernameChange(new_username) {
		this.setState({
			username: new_username,
			askForNewUsername: false,
			Room: this.state.Room,
		});
	}
	handleRoomChange(Room) {
		this.setState({
			...this.state,
			Room: Room,
		});
	}
	render() {
		return (
			<div>
				{this.state.askForNewUsername ? (
					<UsernameComponent
						handleUsernameChange={this.handleUsernameChange}
					/>
				) : null}
				<RoomsComponent handleRoomChange={this.handleRoomChange} />
				<ChatComponent
					username={this.state.username}
					Room={this.state.Room}
				/>
			</div>
		);
	}
}

export default App;
