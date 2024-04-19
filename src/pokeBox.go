package main

import (
	"errors"

	"github.com/google/uuid"
	"github.com/janmmiranda/repl-dex/internal/pokeapi"
)

type pokeBox struct {
	box       map[uuid.UUID]pokeapi.Wpokemon
	caughtMap map[string][]uuid.UUID
}

func pokeBoxCreate() pokeBox {
	return pokeBox{
		box:       make(map[uuid.UUID]pokeapi.Wpokemon),
		caughtMap: make(map[string][]uuid.UUID),
	}
}

func (pb *pokeBox) pokemonAdd(pokemon pokeapi.Pokemon) {
	id := uuid.New()
	wp := pokeapi.NewWpokemon(id, pokemon)
	pb.box[id] = wp
	if val, ok := pb.caughtMap[pokemon.Name]; ok {
		val = append(val, id)
	} else {
		tempSlice := make([]uuid.UUID, 0)
		tempSlice = append(tempSlice, id)
		pb.caughtMap[pokemon.Name] = tempSlice
	}
}

func (pb *pokeBox) pokemonGet(pokemonName string) (pokeapi.Pokemon, error) {
	if val, ok := pb.caughtMap[pokemonName]; ok {
		tempId := val[0]
		wpokemonRes := pb.box[tempId]
		pokemonRes := wpokemonRes.Pokemon
		return pokemonRes, nil
	}
	return pokeapi.Pokemon{}, errors.New("pokemon not caught")
}
