package importer

import (
	"github.com/dohr-michael/graphdb-presentation/pokeapi"
)

func LoadItemCategories() interface{} {
	var result []*pokeapi.ItemCategory
	translationLoader(loaderParams{
		BaseUrl:              "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/item_categories.csv",
		TranslationUrl:       "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/item_category_prose.csv",
		IdentifierField:      "identifier",
		IdField:              "id",
		TranslationIdField:   "item_category_id",
		TranslationLangField: "local_language_id",
		TranslationField:     "name",
		Factory: func(id int, identifier string) pokeapi.WithIdentity {
			return &pokeapi.ItemCategory{
				Identity: pokeapi.NewIdentity(id, identifier),
			}
		},
	}, func(item pokeapi.WithIdentity, row map[string]string) {
		i := item.(*pokeapi.ItemCategory)
		i.PocketId = parseId(row["pocket_id"])
		result = append(result, i)
	})

	return result
}
