package importer

import "github.com/dohr-michael/graphdb-presentation/pokeapi"

type ImportResult struct {
	Generations       []*pokeapi.Generation
	Items             []*pokeapi.Item
	ItemCategories    []*pokeapi.ItemCategory
	Pockets           []*pokeapi.Pocket
	Types             []*pokeapi.Type
	Moves             []*pokeapi.Move
	MoveMethods       []*pokeapi.MoveMethod
	EvolutionTriggers []*pokeapi.EvolutionTrigger
	Pokemons          []*pokeapi.Pokemon
}

func All() interface{} {
	fn := func(call func() interface{}, apply func(r interface{})) <-chan int {
		res := make(chan int, 1)
		go func() {
			apply(call())
			res <- 1
		}()
		return res
	}

	result := &ImportResult{}

	generations := fn(LoadGenerations, func(r interface{}) { result.Generations = r.([]*pokeapi.Generation) })
	itemCategories := fn(LoadItemCategories, func(r interface{}) { result.ItemCategories = r.([]*pokeapi.ItemCategory) })
	items := fn(LoadItems, func(r interface{}) { result.Items = r.([]*pokeapi.Item) })
	pockets := fn(LoadPockets, func(r interface{}) { result.Pockets = r.([]*pokeapi.Pocket) })
	types := fn(LoadTypes, func(r interface{}) { result.Types = r.([]*pokeapi.Type) })
	moves := fn(LoadMoves, func(r interface{}) { result.Moves = r.([]*pokeapi.Move) })
	moveMethods := fn(LoadMoveMethods, func(r interface{}) { result.MoveMethods = r.([]*pokeapi.MoveMethod) })
	evolutionTriggers := fn(LoadEvolutionTriggers, func(r interface{}) { result.EvolutionTriggers = r.([]*pokeapi.EvolutionTrigger) })
	pokemons := fn(LoadPokemons, func(r interface{}) { result.Pokemons = r.([]*pokeapi.Pokemon) })

	<-generations
	<-itemCategories
	<-items
	<-pockets
	<-types
	<-moves
	<-moveMethods
	<-evolutionTriggers
	<-pokemons

	return result
}
