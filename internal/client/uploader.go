package client

import(
	"io"
	"log"
	"net"
	"os"
	"encoding/json"
	"github.com/saketh-exe/shareGO/internal"
)

func sendFile(fileData FileMetadata){
	log.Println("sending File")

	receiver := "localhost:" + internal.PORT
	conn,err := net.Dial("tcp",receiver)
	
	if err != nil{log.Fatalf("ERROR connection to peer %v \n",err);return}
	
	defer conn.Close()

	fileDataJson,err := json.Marshal(fileData)
	if err != nil {
		log.Printf("Couldn't convert %v into json",fileData)
	}
	conn.Write(fileDataJson)
	conn.Write([]byte("\n"))
	file,err :=  os.Open(fileData.Path)
	if err != nil{
		log.Println("Can't open the path")
		return
	}

	defer file.Close()
	log.Printf("Steaming : %s\n",fileData.FileName)
	buf := make([]byte, internal.BufferSize) 
	_, err = io.CopyBuffer(conn,file, buf)
	if err != nil{
		log.Printf("ERROR sending file data: %v \n", err)
		return
	}
	log.Printf("FILE: %s SENT SUCCESSFULLY",fileData.FileName)

}