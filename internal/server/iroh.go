package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	iroh "git.coopcloud.tech/decentral1se/iroh-go"
)

var alpn = []byte("goSend/1")

// ConnectIroh dials the remote peer identified by targetNodeID
func ConnectIroh(targetNodeID string) {
	log.Println("Starting Iroh endpoint...")

	preset := iroh.PresetN0()
	opts := iroh.EndpointOptions{
		Preset: &preset,
		Alpns:  &[][]byte{alpn},
	}

	endpoint, err := iroh.EndpointBind(opts)
	if err != nil {
		log.Fatalln("Error creating endpoint:", err)
	}
	defer endpoint.Close()

	endpoint.Online()
	log.Println("Local Endpoint ID:", endpoint.Id())

	// 1. Parse target Node ID string
	remoteID, err := iroh.EndpointIdFromString(targetNodeID)
	if err != nil {
		log.Fatalln("Invalid remote NodeID:", err)
	}

	// 2. Connect to the peer via relay/P2P
	log.Printf("Connecting to peer: %s\n", targetNodeID)
	addr := iroh.NewEndpointAddr(remoteID, nil, nil)
	conn, err := endpoint.Connect(addr, alpn)
	if err != nil {
		log.Fatalln("Error connecting to peer:", err)
	}
	log.Println("Connected to peer successfully!")

	// 3. Open a bidirectional stream
	stream, err := conn.OpenBi()
	if err != nil {
		log.Fatalln("Error opening BiStream:", err)
	}

	// 4. Send local OS information
	err = stream.Send().WriteAll([]byte(runtime.GOOS))
	if err != nil {
		log.Fatalln("Error sending data:", err)
	}
	log.Println("Sent OS:", runtime.GOOS)

	// 5. Read peer's response frame (up to 256 bytes)
	recv := stream.Recv()
	frame, err := recv.Read(256)
	if err != nil {
		log.Fatalln("Error reading from stream:", err)
	}

	fmt.Printf("Received response from peer: %s\n", string(frame))

	// Keep alive until Ctrl+C
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}