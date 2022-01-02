#!/bin/bash

SERVER_PATH="$HOME/go/hunter-codefest-2022-hchat/server/hChatServer.go"
NGROK_PATH="$HOME/ngrok"
PORT=8888
if [[ -f "$SERVER_PATH" ]] && [[ -f "$NGROK_PATH" ]]; then
	 nohup go run $SERVER_PATH > /dev/null &
	 echo "SERVER: $!" > ps.txt
	 nohup $NGROK_PATH tcp $PORT > /dev/null &
	 echo "NGROK: $!" >> ps.txt 
	 echo "Started Procceses"
	 ps=`cat ps.txt`
	 echo -e "\e[1;34m$ps\e[1;m"
	 echo -e "\e[1;31mPlease run command alias ngrokaddr for address\e[1;m"
else
	if [[ ! -f "$SERVER_PATH" ]]; then
		 echo "Server path is wrong"
	 fi
	 if [[ ! -f "$NGROK_PATH" ]]; then
		 echo "Ngrok path is wrong"
	 fi
fi

