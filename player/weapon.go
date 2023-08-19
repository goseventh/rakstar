package player

import (
	"errors"
	"github.com/goseventh/rakstar/internal/natives"
	"math/rand"
)

type Weapon struct {
	ID   int
	Ammo int
}

var (
	ErrFailSelectWeapon = errors.New("Failure to select player's weapon.")
)

// Invocar esta função selecionará a arma que o jogador está
// segurando. O ID da arma é a posição das armas que o jogador
// possuí atualmete.
//
// Consulte:
// - https://pkg.go.dev/github.com/goseventh/rakstar/player#PlayerBuilder.AddWeapon
func (pb *PlayerBuilder) SelectWeapon(weapon int) error {
	if weapon > len(pb.ListWeapons) || weapon < 0 {
		return ErrFailSelectWeapon
	}

	w := pb.ListWeapons[weapon]
	sucess := natives.SetPlayerArmedWeapon(pb.ID, w.ID)

	if !sucess {
		return ErrFailSelectWeapon
	}
	return nil
}

// Invocar esta função selecionará aleatóriamente a arma que o jogador
// está segurando. A arma selecionada será uma das armas presentes na
// lista de armas
//
// Consulte:
// - https://pkg.go.dev/github.com/goseventh/rakstar/player#PlayerBuilder.AddWeapon
func (pb *PlayerBuilder) SelectRandomWeapon() error {
	var list []Weapon
	var weapon Weapon

	copy(list, pb.ListWeapons)
	rand.Shuffle(len(list), func(idx, weapon int) {
		weapon = weapon
	})

  sucess := natives.SetPlayerArmedWeapon(pb.ID, weapon.ID)
	if !sucess {
		return ErrFailSelectWeapon
	}
	return nil
}

// Invocar esta função adicionará a arma e sua munição recebidas
// no parametro, à lista de armas do jogador
func (pb *PlayerBuilder) AddWeapon(weapon, ammo int) {
	if ammo < 0 {
		ammo = 30
	}
	w := Weapon{
		ID:   weapon,
		Ammo: ammo,
	}

	pb.ListWeapons = append(pb.ListWeapons, w)
}

// Invocar esta função limpará a lista de armas do jogador
func (pb *PlayerBuilder) ClearListWeapons() {
	pb.ListWeapons = nil
}
