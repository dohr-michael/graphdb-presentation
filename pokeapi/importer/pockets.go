package importer

import (
	"github.com/dohr-michael/graphdb-presentation/pokeapi"
)

func LoadPockets() interface{} {
	var result []*pokeapi.Pocket
	translationLoader(loaderParams{
		BaseUrl:              "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/item_pockets.csv",
		TranslationUrl:       "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/item_pocket_names.csv",
		IdentifierField:      "identifier",
		IdField:              "id",
		TranslationIdField:   "item_pocket_id",
		TranslationLangField: "local_language_id",
		TranslationField:     "name",
		Factory: func(id int, identifier string) pokeapi.WithIdentity {
			return &pokeapi.Pocket{
				Identity: pokeapi.NewIdentity(id, identifier),
			}
		},
	}, func(item pokeapi.WithIdentity, row map[string]string) {
		result = append(result, item.(*pokeapi.Pocket))
	})

	return result
}
