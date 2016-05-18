package radius_go

import (
	"fmt"
	"net"
	"os"
)


var progName = "server"
var host = "0.0.0.0"


func ServerStartUp() error {
	addr, err := net.ResolveUDPAddr("udp", host+":"+"1812")
	if err != nil {
		fmt.Fprintf(os.Stderr, "[%s][%s]\t%s\n", progName, "ERROR", "Binding Address Failed")
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[%s][%s]\t%s\n", progName, "ERROR", "Listening Failed")
		return err
	}else {
		fmt.Fprintf(os.Stdout, "[%s][%s]\t%s\n", progName, "NORMAL", "Listening...")
	}

	defer conn.Close()

	for {
		err := handleRequest(conn)
		if err != nil{
			return err
		}
	}
}

func handleRequest(conn *net.UDPConn) error {
	inBuff := make([]byte, 1024)
	_, remoteAddr, err := conn.ReadFromUDP(inBuff)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[%s][%s] %s\n", progName, "WARNING", "UDP reading Error")
		return err
	}
	fmt.Printf("Messege is: "+string(inBuff))
	conn.WriteToUDP([]byte("Got it"), remoteAddr)
	return nil
}
