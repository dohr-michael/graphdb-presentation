package importer

import (
	"github.com/dohr-michael/graphdb-presentation/pokeapi"
)

func LoadMoves() interface{} {
	var result []*pokeapi.Move
	translationLoader(loaderParams{
		BaseUrl:              "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/moves.csv",
		TranslationUrl:       "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/move_names.csv",
		IdentifierField:      "identifier",
		IdField:              "id",
		TranslationIdField:   "move_id",
		TranslationLangField: "local_language_id",
		TranslationField:     "name",
		Factory: func(id int, identifier string) pokeapi.WithIdentity {
			return &pokeapi.Move{
				Identity: pokeapi.NewIdentity(id, identifier),
			}
		},
	}, func(item pokeapi.WithIdentity, row map[string]string) {
		i := item.(*pokeapi.Move)
		i.Generation = parseId(row["generation_id"])
		i.Type = parseId(row["type_id"])
		i.Power = parseInt(row["power"])
		i.Pp = parseInt(row["pp"])
		i.Accuracy = parseInt(row["accuracy"])
		result = append(result, i)
	})

	return result
}
