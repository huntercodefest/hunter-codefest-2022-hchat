import React from "react";
import ReactDOM from "react-dom";
import App from "./components/App"

// Render app

// class Serve extends React.Component{
    
//     constructor(){
//         super()
//         this.state = {
//             data: null
//         }
//     }
//     componentDidMount() {
//         this.callBackendAPI()
//           .then(res => this.setState({ data: res.express }))
//           .catch(err => console.log(err));
//       }
//         // fetching the GET route from the Express server which matches the GET route from server.js
//       callBackendAPI = async () => {
//         const response = await fetch('/express_backend');
//         const body = await response.json();
    
//         if (response.status !== 200) {
//           throw Error(body.message) 
//         }
//         return body;
//       };
//       render(){
//         return(
//             <App />
//         )
//       }
// }

ReactDOM.render(<App />, document.getElementById("root"));
