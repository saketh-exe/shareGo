package main

import (
	"fmt"
	"os"

	"github.com/saketh-exe/shareGO/internal/client"
	_ "github.com/saketh-exe/shareGO/internal/client"
)

func main(){
	args := os.Args
	path := args[0]
	command := args[1]
	filepaths := args[2:]
	
	fmt.Printf("Script path : %s; command : %s;\n", path, command)
	for _,path := range filepaths{
		fmt.Printf("%s \n", path)
	}

	client.UploadFile(filepaths)
	
}