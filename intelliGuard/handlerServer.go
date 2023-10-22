package intelliGuard

import "net"
import "log"

// esta é uma função interna para a manipulação dos dados do
// servidor proxy - este que exportará as conexões do servidor samp,
// normalmente na porta 7777.
func handlerServerData(receiverChan <-chan interface{}, writerChan chan<- interface{}, ln *net.UDPConn) {
	defer ln.Close()
	buffer := make([]byte, 1500)
	go func() {
		for {
			n, err := ln.Read(buffer)
			if err != nil {
				log.Println(err)
				continue
			}
			writerChan <- buffer[:n]
		}
	}()

	for data := range receiverChan {
    log.Println("handler server enviou ao cliente samp")
		ln.Write(data.([]byte))
	}
}
