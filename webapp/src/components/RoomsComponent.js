import React from "react";
import BackButton from "./BackButton";
import SearchBar from "./SearchBar";
import RoomsContainer from "./RoomsContainer";
import classes from "../css/rooms.module.css";
import logo from "../hchatlogo.png";
import { Room, RoomList, Schools } from "./Rooms";
class RoomsComponent extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			searchList: Schools,
			searchValue: "",
			displayList: Schools.CurrList,
		};
		// Function binds
		this.handleListChange = this.handleListChange.bind(this);
		this.handleBackClick = this.handleBackClick.bind(this);
		this.handleSearchMod = this.handleSearchMod.bind(this);
	}
	handleSearchMod(event) {
		const searchVal = event.target.value;
		// modify displaylist
		// change value
		this.setState({
			...this.state,
			searchValue: searchVal,
			displayList: this.state.searchList.CurrList.filter((room) =>
				room.room_desc.toLowerCase().includes(searchVal.toLowerCase())
			),
		});
	}
	handleBackClick() {
		// modify searchlist
		// If clicked back in school list
		if (!this.state.searchList.PrevList) {
			return;
		}
		// otherwise set searchList and displayList to prevlist
		this.setState({
			...this.state,
			searchList: this.state.searchList.PrevList,
			displayList: this.state.searchList.PrevList.CurrList,
			searchValue: "",
		});
		this.props.updateListDesc(this.state.searchList.PrevList.listDesc);
	}
	handleListChange(room) {
		// load in paths from school folder
		const rel_path = `./CUNY-CLASSES-JSON/${room.school}`;
		const paths = require(`${rel_path}/_PATHS.json`);
		let roomList = null;
		// if clicked on a school
		if (room.major === null) {
			// create new roomlist of school, major
			roomList = new RoomList(
				paths.map(
					(path) =>
						new Room(room.school, path.major, null, path.major)
				),
				this.state.searchList,
				room.room_desc
			);
		}
		// if clicked on a major
		else if (room.room_num === null) {
			// return the path of the json file with school
			const p = paths.filter((path) => path.major === room.major);
			// return the json file of classes
			const classes = require(`${rel_path}/${p[0].path}`);
			// create new roomlist of school, major, room number, room desc
			roomList = new RoomList(
				classes.map(
					(_class) =>
						new Room(
							_class.school,
							_class.major,
							_class.room_num,
							_class.room_desc
						)
				),
				this.state.searchList,
				room.room_desc
			);
		}
		// if clicked on a class
		// handle in room change
		else {
			return;
		}
		this.setState({
			searchValue: "",
			searchList: roomList,
			displayList: roomList.CurrList,
		});
	}
	render() {
		return (
			<div className={classes.roomsComponent}>
				<div className={classes.header}>
					<span className={classes.desc}>
						<BackButton handleBackClick={this.handleBackClick} />
						<p>{this.props.listDesc}</p>
						<SearchBar
							searchList={this.state.searchList.CurrList}
							value={this.state.searchValue}
							handleSearchMod={this.handleSearchMod}
						/>
					</span>
					<div className={classes.search}>
						<img
							src={logo}
							className={classes.logo}
							onClick={() => window.location.href = "http://hchat.org"}
							alt="hchatlogo"></img>
						<p>Found {this.state.displayList.length} rooms</p>
					</div>
				</div>
				<RoomsContainer
					displayList={this.state.displayList}
					handleRoomChange={this.props.handleRoomChange}
					handleListChange={this.handleListChange}
				/>
			</div>
		);
	}
}

export default RoomsComponent;
