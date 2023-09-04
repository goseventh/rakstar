package intelliGuard

import "net"

func CreateReverseProxy() {
  addr , err := net.ResolveUDPAddr("udp", ":7777")
  if err != nil{
    panic(err)
  }
  ln, err := net.ListenUDP("udp", addr)
  if err != nil{
    panic(err)
  }
  go HandlerConn(ln) 
}
