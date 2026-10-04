package client

import(
	"log"
	"os"
	iroh "git.coopcloud.tech/decentral1se/iroh-go"
	"os/signal"
	"syscall"
	"runtime"
)


var alpn = []byte("goSend/1")

func ConnectIroh(){

	log.Println("Starting IroH connection....")

	preset := iroh.PresetN0()
	opts := iroh.EndpointOptions{
			Preset: &preset,
			Alpns: &[][]byte{alpn},
	}

	endpoint , err := iroh.EndpointBind(opts)
	if err != nil {
		log.Fatalln("Error creating an endpoint .... ",err)
	}
	defer endpoint.Close()
	endpoint.Online()
	
	log.Println("Listening with NodeID : ",endpoint.Id())

	go func() {
		for {
			incoming_conn := *endpoint.AcceptNext()
			if incoming_conn == nil {
				// Endpoint was closed or shut down; exit the loop cleanly
				return
			}
			go handleIncoming(*incoming_conn)
		}
	}()

	// Keep alive until Ctrl+C
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}


func handleIncoming(incomming iroh.Incoming){
	accepting, err := incomming.Accept()
	if err != nil {log.Fatalln("Error accepting Incomming connection ",err)}
	
	conn,conn_err := accepting.Connect()
	
	if conn_err != nil {log.Fatalln("Error connecting to Incomming connection ",err)}

	stream,err := conn.AcceptBi()

	if err != nil {log.Fatalln("Error getting steam from Incomming connection ",err)}

	// 1. Send your OS string to the peer
	 stream.Send().WriteAll([]byte(runtime.GOOS))
	
	// 2. Read the peer's response (e.g. up to 256 bytes)
	recv := stream.Recv()
	frame, err := recv.Read(256)
	if err != nil {
		panic(err)
	}

	println(string(frame))
}

