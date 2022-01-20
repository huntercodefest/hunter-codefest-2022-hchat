class ChatComponent extends React.Component {
    constructor(props) {
		// connect to server here
		super(props);
		this.ws = null;
		this.state = {
			messages: [],
			inputvalue: "",
		};
	}
	connectToServer() {
		// TODO
		this.ws.onopen = () => {
			this.ws.send(
                `#${this.props.Room.room_num}
                _${this.props.username}
                :`
                );
		};
		this.ws.onmessage = (msg) => {
			this.receiveFromServer(message);
		};
	}
	sendToServer(msg) {
		if (this.validateMessage(msg)){
            this.ws.send(this.processMessage(msg));
        }
	}
	receiveFromServer(rcv_msg) {
		// TODO
		// Parse received message
		// And make new message component
		// with time and username and message
        const user, msg;
        for (let i = 0; i < rcv_msg.length; i++){
            const user_end_index = rcv_msg.indexOf(':')
            user = rcv_msg.substring(
                rcv_msg.indexOf('_')+1, 
                user_end_index)
            msg = rcv_msg.substring(
                user_end_index+1)
        }

        const time = new Date();
        const message = <Message time={time} username={user} message={msg}/>;
		this.setState({
			...this.state,
			messages: this.state.messages.concat(message),
		});
	}
    validateMessage(msg) {
		// placeholder
		return msg.length > 0 && msg.length <= 500;
	}
    processMessage(msg){
        return `#${this.props.Room.room_num}
        _${this.props.username}
        :${msg}`
    }
	handleSubmit(event) {
		event.preventDefault();
		this.sendToServer(event.target[0].value);
		// Send message to server
	}
    handleChange(event){
        this.setState({
			...this.state,
			inputvalue: event.target.value
        })
    }
    attemptConnection(){
        // If ws connection already exists or username is empty
        if (this.ws !== null || this.props.username === ""){
            return
        }
        this.ws = new WebSocket()
        this.connectToServer()

    }
	render() {
        this.attemptConnection()
		return (
			<div>
				<ChatRoom
					Room={this.props.Room_Desc}
					Messages={this.state.messages}
				/>
				<MessageInput
					handleSubmit={this.handleSubmit}
                    handleChange={this.handleChange}
					value={this.state.inputvalue}
				/>
			</div>
		);
	}
}

export default ChatComponent