package player

import (
	"time"

	"github.com/goseventh/rakstar/internal/natives"
)
type wallet struct{
  balance int
}


// Invocar esta função atualizará o dinheiro do jogador,
// diferentemente da native SAMP, esta função não acrescenta
// dinheiro ao jogador mas seta a quantia. Quando invocada, 
// setará a quantidade de dinheiro e em seguida, a função
// interna do RakStar será responsável em manter a label
// do jogo - esta que informa a quantia em dinheiro para o
// jogador. Se alguma manipulação externa, como por exemplo,
// o uso de sobeits, adulterar maliciosamente a label do jogador,
// o RakStar automaticamente corrigirá o valor, setando-o de volta
// ao valor original
func(pb*PlayerBuilder) SetBalance(quanty int) {
  pb.balance = quanty
}

// Invocar esta função retornará a quantia de dinheiro
// do jogador.
func (pb*PlayerBuilder) GetBalance() int{
  return pb.balance
}

// Esta função é uma função interna do RakStar, sua
// utilização deve ser apenas para o framework. A
// função atualiza a label de dinheiro do jogador.
func (pb *PlayerBuilder) updateMoney(quanty int) bool {
	money := natives.GetPlayerMoney(pb.ID)
	if money != quanty {
		natives.ResetPlayerMoney(pb.ID)
		time.Sleep(time.Second)
		sucess := natives.GivePlayerMoney(pb.ID, quanty)
		return sucess
	}
  return true
}
