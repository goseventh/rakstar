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

// SelectWeapon seleciona a arma que o jogador está segurando.
// Ele recebe um argumento inteiro 'weapon' que representa o ID da arma. 
// O ID da arma é a posição das armas que o jogador possui atualmente.
// Se o ID da arma for maior do que o número de armas que o jogador possui 
// ou menor do que zero, o método retorna um erro ErrFailSelectWeapon.
// Caso contrário, ele obtém a arma correspondente da lista de armas do jogador 
// e chama a função nativa SetPlayerArmedWeapon com o ID do jogador e o ID da arma.
// Se a função nativa retornar false, indicando que a arma não pôde ser selecionada, 
// o método retorna um erro ErrFailSelectWeapon.
// Se a arma for selecionada com sucesso, o método retorna nil, indicando que não 
// houve erros.
// Consulte também: [AddWeapon]
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

// SelectRandomWeapon seleciona aleatoriamente uma arma da lista de armas do jogador.
// O método cria uma cópia da lista de armas do jogador e embaralha-a.
// Em seguida, ele seleciona a última arma na lista embaralhada.
// O método chama a função nativa SetPlayerArmedWeapon com o ID do jogador e o ID 
// da arma selecionada.
// Se a função nativa retornar false, indicando que a arma não pôde ser selecionada, 
// o método retorna um erro ErrFailSelectWeapon.
// Se a arma for selecionada com sucesso, o método retorna nil, indicando que não 
// houve erros.
func (pb *PlayerBuilder) SelectRandomWeapon() error {
	var list []Weapon

	copy(list, pb.ListWeapons)
	rand.Shuffle(len(list), func(i, j int) {
    list[i], list[j] = list[j], list[i]
	})

  weapon := list[len(list)-1]
  sucess := natives.SetPlayerArmedWeapon(pb.ID, weapon.ID)
	if !sucess {
		return ErrFailSelectWeapon
	}
	return nil
}

// AddWeapon adicionará na lista de armas do jogador, a respectiva arma e sua munição
// informadas no parâmetro. Se a munição for menor que zero será setado o valor padrão
// 30. Informar um ID de arma inválido ocasionará em falhas nas funções seletivas 
// [SelectWeapon] e [SelectRandomWeapon]
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

// ClearListWeapons removerá a lista de armas do jogador
func (pb *PlayerBuilder) ClearListWeapons() {
	pb.ListWeapons = nil
}
