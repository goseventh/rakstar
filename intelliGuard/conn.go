package intelliGuard

import "net"

// Esta função abre uma proxy reversa para exportar os 
// pacotes do servidor samp com proteção de ataques de 
// negação de serviço (DDOS) e reconstrução de pacotes
// danificados.
// Exemplo: CreateReverseProxy("localhost:3000")
func CreateReverseProxy(source string) {
	if source == "" {
		source = "localhost:3000"
	}

	chanServer := createServer()
	chanSamp := createClientSamp(source)
	go linkServerToSamp(chanServer, chanSamp)
}

// esta é uma função interna da proxy reversa para linkar a proxy ao servidor samp
func linkServerToSamp(chanServer, chanSamp chan interface{}) {
	for {
		select {
		case data := <-chanServer:
			chanSamp <- data
		case data := <-chanSamp:
			chanServer <- data
		}
	}
}

// esta é uma função interna da proxy reversa para criar um servidor UDP
// que exportará os pacotes do servidor samp
func createServer() chan interface{} {
	addr, err := net.ResolveUDPAddr("udp", ":7777")
	if err != nil {
		panic(err)
	}
	ln, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	channel := make(chan interface{})
	go handlerServerData(channel, ln)
	return channel
}

// esta é uma função interna da proxy reversa para criar um cliente UDP que
// se conectará com o servidor SAMP para portar os pacotes
// do servidor UDP da proxy para o servidor do samp.
func createClientSamp(addr string) chan interface{} {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		panic(err)
	}
	channel := make(chan interface{})
	go handlerClientData(channel, conn)
	return channel
}
