package intelliGuard

import "net"
import "log"
import "fmt"

func ConnectToSamp(addr string) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		panic(err)
	}
	go HandlerSamp(conn)
}

func HandlerSamp(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1200)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("package: %v", buffer[:n])
	}
}
