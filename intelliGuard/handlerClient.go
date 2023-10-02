package intelliGuard

import "net"
import "log"
import "fmt"
import "slices"

// esta é uma função interna da proxy reversa para a manipulação dos dados
// do cliente que se conecta ao servidor samp - este que por padrão deve
// criar uma conexão na porta 3000.
func handlerClientData(receiverChan <-chan interface{}, writerChan chan<- interface{}, conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1200)
	var lastPackage []byte
	go func() {
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				panic(err)
				log.Println(err)
			}
      fmt.Printf("package: %v\n", buffer[:n])
      log.Println("handler client enviou dados para handler server")
			writerChan <- buffer[:n]
			// lastPackage = buffer[:n]
		}

	}()

	for data := range receiverChan {
		if slices.Equal(data.([]byte), lastPackage) {
			panic("recebeu o dado que ele mesmo enviou")
		}
		log.Println("handler cliente recebeu dados do handler server")
		conn.Write(data.([]byte))
	}
}
