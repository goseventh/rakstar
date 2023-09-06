package intelliGuard

import "net"
import "log"
import "fmt"

// esta é uma função interna da proxy reversa para a manipulação dos dados
// do cliente que se conecta ao servidor samp - este que por padrão deve
// criar uma conexão na porta 3000. 
func handlerClientData(channel chan interface{}, conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1200)
	go func() {
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				log.Println(err)
			}
      channel <- buffer[:n]
			fmt.Printf("package: %v", buffer[:n])
		}
	}()

  for data := range channel{
    conn.Write(data.([]byte))
  }
}
