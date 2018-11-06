package importer

import (
	"github.com/dohr-michael/graphdb-presentation/pokeapi"
)

func LoadGenerations() interface{} {
	var result []*pokeapi.Generation
	translationLoader(loaderParams{
		BaseUrl:              "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/generations.csv",
		TranslationUrl:       "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/generation_names.csv",
		IdentifierField:      "identifier",
		IdField:              "id",
		TranslationIdField:   "generation_id",
		TranslationLangField: "local_language_id",
		TranslationField:     "name",
		Factory: func(id int, identifier string) pokeapi.WithIdentity {
			return &pokeapi.Generation{
				Identity: pokeapi.NewIdentity(id, identifier),
			}
		},
	}, func(item pokeapi.WithIdentity, row map[string]string) {
		result = append(result, item.(*pokeapi.Generation))
	})

	return result
}
