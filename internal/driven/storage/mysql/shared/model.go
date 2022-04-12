package shared

import (
	"github.com/Haraj-backend/hex-pokebattle/internal/core/entity"
)

type PokeRow struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Health    int    `db:"health"`
	MaxHealth int    `db:"max_health"`
	Attack    int    `db:"attack"`
	Defense   int    `db:"defense"`
	Speed     int    `db:"speed"`
	AvatarURL string `db:"avatar_url"`
}

type PokeRows []PokeRow

func (r *PokeRow) ToPokemon() *entity.Pokemon {
	return &entity.Pokemon{
		ID:   r.ID,
		Name: r.Name,
		BattleStats: entity.BattleStats{
			Health:    r.Health,
			MaxHealth: r.MaxHealth,
			Attack:    r.Attack,
			Defense:   r.Defense,
			Speed:     r.Speed,
		},
		AvatarURL: r.AvatarURL,
	}
}

func (r PokeRows) ToPokemons() []entity.Pokemon {
	var pokemons []entity.Pokemon
	for _, row := range r {
		pokemons = append(pokemons, *row.ToPokemon())
	}
	return pokemons
}

func NewPokeRow(pokemon *entity.Pokemon) *PokeRow {
	return &PokeRow{
		ID:        pokemon.ID,
		Name:      pokemon.Name,
		Health:    pokemon.BattleStats.Health,
		MaxHealth: pokemon.BattleStats.MaxHealth,
		Attack:    pokemon.BattleStats.Attack,
		Defense:   pokemon.BattleStats.Defense,
		Speed:     pokemon.BattleStats.Speed,
		AvatarURL: pokemon.AvatarURL,
	}
}
