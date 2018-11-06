package importer

import (
	"github.com/dohr-michael/graphdb-presentation/pokeapi"
)

func LoadTypes() interface{} {
	var result []*pokeapi.Type
	translationLoader(loaderParams{
		BaseUrl:              "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/types.csv",
		TranslationUrl:       "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/type_names.csv",
		IdentifierField:      "identifier",
		IdField:              "id",
		TranslationIdField:   "type_id",
		TranslationLangField: "local_language_id",
		TranslationField:     "name",
		Factory: func(id int, identifier string) pokeapi.WithIdentity {
			return &pokeapi.Type{
				Identity: pokeapi.NewIdentity(id, identifier),
			}
		},
	}, func(item pokeapi.WithIdentity, row map[string]string) {
		result = append(result, item.(*pokeapi.Type))
	})

	return result
}
