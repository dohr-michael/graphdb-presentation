package importer

import (
	"github.com/dohr-michael/graphdb-presentation/pokeapi"
)

func LoadPokemons() interface{} {
	var result []*pokeapi.Pokemon
	pokeMap := map[pokeapi.Id]*pokeapi.Pokemon{}

	translationLoader(loaderParams{
		BaseUrl:              "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon_species.csv",
		TranslationUrl:       "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon_species_names.csv",
		IdentifierField:      "identifier",
		IdField:              "id",
		TranslationIdField:   "pokemon_species_id",
		TranslationLangField: "local_language_id",
		TranslationField:     "name",
		Factory: func(id int, identifier string) pokeapi.WithIdentity {
			return &pokeapi.Pokemon{
				Identity: pokeapi.NewIdentity(id, identifier),
			}
		},
	}, func(item pokeapi.WithIdentity, row map[string]string) {
		p := item.(*pokeapi.Pokemon)
		p.EvolveInto = []pokeapi.Evolution{}

		pokeMap[p.Id] = p
		p.Generation = parseId(row["generation_id"])
		result = append(result, p)
	})

	downloadAndRead("https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon.csv", func(row row) {
		id := parseId(row["id"])
		if pokeMap[id] != nil {
			pokemon := pokeMap[id]
			pokemon.Height = parseInt(row["height"])
			pokemon.Weight = parseInt(row["weight"])
			pokemon.BaseExp = parseInt(row["base_experience"])
		}
	})

	downloadAndRead("https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon_types.csv", func(row row) {
		id := parseId(row["pokemon_id"])
		if pokeMap[id] != nil {
			pokeMap[id].Types = append(pokeMap[id].Types, parseId(row["type_id"]))
		}
	})

	downloadAndRead("https://raw.githubusercontent.com/PokeAPI/pokeapi/master/data/v2/csv/pokemon_evolution.csv", func(row row) {
		from := parseId(row["id"])
		evolve := pokeapi.Evolution{}
		evolve.Into = parseId(row["evolved_species_id"])
		evolve.TriggerType = parseId(row["evolution_trigger_id"])
		if row["minimum_level"] != "" {
			evolve.MinimumLevel = parseInt(row["minimum_level"])
		}
		if row["trigger_item_id"] != "" {
			evolve.TriggerItem = parseId(row["trigger_item_id"])
		}
		if pokeMap[from] != nil {
			pokeMap[from].EvolveInto = append(pokeMap[from].EvolveInto, evolve)
		}
	})

	downloadAndRead("https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon_moves.csv", func(row row) {
		id := parseId(row["pokemon_id"])
		if pokeMap[id] != nil {
			link := pokeapi.MoveLink{
				MoveId:     parseId(row["move_id"]),
				MoveMethod: parseId(row["pokemon_move_method_id"]),
				Level:      parseInt(row["level"]),
			}
			pokeMap[id].FightWith = append(pokeMap[id].FightWith, link)
		}
	})

	return result
}
