#!/bin/bash

SERVER="go/server/server.go"
PORT=8888
if [[ -f "$HOME/$SERVER" ]] && [[ -f "$HOME/ngrok" ]]; then
	 nohup go run $HOME/$SERVER > /dev/null &
	 echo "SERVER: $!" > ps.txt
	 nohup ./ngrok tcp $PORT > /dev/null &
	 echo "NGROK: $!" >> ps.txt 
	 echo "Started Procceses"
	 ps=`cat ps.txt`
	 echo -e "\e[1;34m$ps\e[1;m"
	 echo -e "\e[1;31mPlease run command alias ngrokaddr for address\e[1;m"
fi

