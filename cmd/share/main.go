package main

import (
	
	"os"

	"github.com/saketh-exe/shareGO/internal/client"
	"github.com/saketh-exe/shareGO/internal/server"
)

func main(){
	args := os.Args
	command := args[1]
	if command == "start"{
		server.StartServer()
	
	}else if command == "iroh" {
		connection := args[2]
		server.ConnectIroh(connection)
	}else{

		client.ConnectIroh()
		
	}
	
}