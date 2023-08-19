package player

import "github.com/goseventh/rakstar/internal/natives"

type Weapon struct {
	ID   int
	Ammo int
}

// Invocar esta função selecionará a arma que o jogador está
// segurando
func (pb *PlayerBuilder) SelectWeapon(weapon int) error {
	sucess := natives.SetPlayerArmedWeapon(pb.ID, weapon)
	if !sucess {
		return FailSelectWeapon
	}
	return nil
}
