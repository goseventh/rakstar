package intelliGuard

import "net"
import "log"

// esta é uma função interna da proxy reversa para a manipulação dos dados do
// servidor proxy - este que exportará as conexões do servidor samp,
// normalmente na porta 7777.
func handlerServerData(channel chan interface{}, ln *net.UDPConn) {
	defer ln.Close()
	buffer := make([]byte, 1200)
	go func() {
		for {
			n, err := ln.Read(buffer)
			if err != nil {
				log.Println(err)
				continue
			}
			channel <- buffer[:n]
		}
	}()

  for data := range channel{
    ln.Write(data.([]byte))
  }
}
