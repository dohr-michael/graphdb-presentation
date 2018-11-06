package importer

import "github.com/dohr-michael/graphdb-presentation/pokeapi"

func LoadMoveMethods() interface{} {
	var result []*pokeapi.MoveMethod
	translationLoader(loaderParams{
		BaseUrl:              "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon_move_methods.csv",
		TranslationUrl:       "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon_move_method_prose.csv",
		IdentifierField:      "identifier",
		IdField:              "id",
		TranslationIdField:   "pokemon_move_method_id",
		TranslationLangField: "local_language_id",
		TranslationField:     "name",
		Factory: func(id int, identifier string) pokeapi.WithIdentity {
			return &pokeapi.MoveMethod{
				Identity: pokeapi.NewIdentity(id, identifier),
			}
		},
	}, func(item pokeapi.WithIdentity, row map[string]string) {
		i := item.(*pokeapi.MoveMethod)
		result = append(result, i)
	})

	return result
}

