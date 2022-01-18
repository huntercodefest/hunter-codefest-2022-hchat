import React from "react"
import classes from "../app.module.css"

class ChatBlock extends React.Component {
    constructor(props){
        super(props);
        this.state = {

        };
        this.message = {
            username: this.props.username,
            message: this.props.message
        }
        console.log(this.message.username + ":", this.message.username);
    }

    render(){
		return (
			<div className={classes.message} style={{ width: "100%" }}>
				<p>{this.message.username}: {this.message.message}</p>
			</div>
		);
	};
}

export default ChatBlock