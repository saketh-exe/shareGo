package main

import (
	"fmt"
	"os"

	"github.com/saketh-exe/shareGO/internal/client"
	"github.com/saketh-exe/shareGO/internal/server"
)

func main(){
	args := os.Args
	path := args[0]
	command := args[1]
	if command == "start"{
		server.StartServer()
	
	}else{

		filepaths := args[2:]
		fmt.Printf("Script path : %s; command : %s;\n", path, command)
		for _,path := range filepaths{
			fmt.Printf("%s \n", path)
		}
		
		client.UploadFile(filepaths)
	}
	
}