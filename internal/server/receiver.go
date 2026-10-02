package server

import (
	"log"
	"net"
	"os"
	"bufio"
	"encoding/json"
	"io"
	_"strings"
	"github.com/joho/godotenv"
)



func StartServer(){
	
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error Loading ENV")
	}
	portNo := os.Getenv("PORT")	
	PORT := ":" + portNo
	if PORT == ""{
		log.Fatal("Port Not Specified")
	}
	listener,err := net.Listen("tcp",PORT)
	if err != nil{
		log.Fatalf("ERROR can't listen on PORT %s, ERROR : %v",PORT,err)
	}
	defer listener.Close()
	log.Println("Started FILE SERVER")
	for {
		conn,err := listener.Accept()
		if err != nil {
			log.Printf("ERROR connection to peer %v\n",err)
		}
		go fileReciver(conn)
	}
	
}

func fileReciver(conn net.Conn){
	defer conn.Close()
	peerAddr := conn.RemoteAddr().String()
	log.Printf("New peer connected %s\n",peerAddr)

	// get the file metadata first 

	reader := bufio.NewReader(conn)

	rawfileMetadata,err := reader.ReadString('\n')
	if err != nil{log.Println("Error reading metadata")}

	var fileMetadata map[string]any
	json.Unmarshal([]byte(rawfileMetadata),&fileMetadata)
	fileName := fileMetadata["FileName"].(string)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {log.Println("Error Creating new File")}	
	defer file.Close()

	io.Copy(file,reader)	
	
	
	}


