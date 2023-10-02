package intelliGuard

import "net"
import "log"

// Esta função abre uma proxy reversa para exportar os
// pacotes do servidor samp com proteção de ataques de
// negação de serviço (DDOS) e reconstrução de pacotes
// danificados.
// 
// Recebe-se Source e Destino;
// Exemplo: CreateReverseProxy("localhost:3000", ":7777")
//
// Se uma string vazia for fornecida, a source padrão será
// "localhost:3000" e dst será "7777"
func CreateReverseProxy(source, dst string) {
	if source == "" {
		source = "localhost:3000"
	}
  if dst == ""{
    dst = ":7777"
  }

	receiverServer, writerServer := createServer(dst)
	receiverClient, writerClient := createClientSamp(source)
	go forwarding(receiverServer, receiverClient, writerServer, writerClient)
}

// esta é uma função interna da proxy reversa para linkar a proxy ao servidor samp
func forwarding(receiverServer, receiverClient chan<- interface{}, writerServer, writerClient <-chan interface{}) {
	log.Println("linker da proxy reversa foi stardado")
	for {
		select {
		case data := <-writerServer:
			log.Println("proxy recebeu dados do jogo cliente")
			receiverClient <- data
		case data2 := <-writerClient:
			log.Println("handler cliente recebeu dados do servidor samp")
			receiverServer <- data2
		}
	}
}

// esta é uma função interna da proxy reversa para criar um servidor UDP
// que exportará os pacotes do servidor samp
func createServer(dst string) (chan<- interface{}, <-chan interface{}) {
	addr, err := net.ResolveUDPAddr("udp", dst)
	if err != nil {
		panic(err)
	}
	ln, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	receiverChan := make(chan interface{})
	writerChan := make(chan interface{})
	go handlerServerData(receiverChan, writerChan, ln)
	log.Println("servidor da proxy reversa foi startado")
	return receiverChan, writerChan
}

// esta é uma função interna da proxy reversa para criar um cliente UDP que
// se conectará com o servidor SAMP para portar os pacotes
// do servidor UDP da proxy para o servidor do samp.
func createClientSamp(addr string) (chan<- interface{}, <-chan interface{}) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		panic(err)
	}

	receiverChan := make(chan interface{})
	writerChan := make(chan interface{})
	go handlerClientData(receiverChan, writerChan, conn)
	log.Println("cliente da proxy reversa foi conectado:", addr)
	return receiverChan, writerChan
}
