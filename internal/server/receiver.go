package server

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	_ "strings"
	"github.com/saketh-exe/shareGO/internal"
)


type FileMetadata struct {
	FileName string
	FileSize int64 // fileStats.Size() is int64
	Path     string
	IsDir    bool
}
func StartServer(){

	portNo := internal.PORT
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

	reader := bufio.NewReaderSize(conn,internal.BufferSize)

	rawfileMetadata,err := reader.ReadString('\n')
	if err != nil{log.Println("Error reading metadata")}

	var fileMetadata FileMetadata
	json.Unmarshal([]byte(rawfileMetadata),&fileMetadata)
	
	fileName := fileMetadata.FileName
	fileSize := fileMetadata.FileSize
	sizeInGB := float64(fileSize) / 1_073_741_824
	log.Printf("Receiving File %s , Size : %.2f GB",fileName,sizeInGB )
	
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {log.Println("Error Creating new File")}	
	defer file.Close()
	
	buf := make([]byte, internal.BufferSize)

// 3. Copy using the custom buffer while keeping the strict size limit
	written , err := io.CopyBuffer(file, reader, buf)
	if err != nil || written != fileSize { // Transmission has stopped
		file.Close()
		os.Remove(file.Name())
		log.Printf("Transmission Failed for file %s , Deleting Partial File....",fileName)
		return
	}
	
	log.Printf("Received File %s , Written Size : %.2f GB",fileName,float64(written)/1_073_741_824 )
	
	}


