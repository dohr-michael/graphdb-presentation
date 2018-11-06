package importer

import "github.com/dohr-michael/graphdb-presentation/pokeapi"

func LoadItems() interface{} {
	var result []*pokeapi.Item
	dic := make(map[pokeapi.Id]*pokeapi.Item)
	translationLoader(loaderParams{
		BaseUrl:              "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/items.csv",
		TranslationUrl:       "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/item_names.csv",
		IdentifierField:      "identifier",
		IdField:              "id",
		TranslationIdField:   "item_id",
		TranslationLangField: "local_language_id",
		TranslationField:     "name",
		Factory: func(id int, identifier string) pokeapi.WithIdentity {
			return &pokeapi.Item{
				Identity: pokeapi.NewIdentity(id, identifier),
				Details:  make(pokeapi.Labels),
			}
		},
	}, func(item pokeapi.WithIdentity, row map[string]string) {
		i := item.(*pokeapi.Item)
		dic[i.Id] = i
		i.Category = parseId(row["category_id"])
		i.Cost = parseInt(row["cost"])
		result = append(result, i)
	})
	downloadAndRead("https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/item_prose.csv", func(row row) {
		id := parseId(row["item_id"])
		lang := parseInt(row["local_language_id"])
		desc := parseString(row["short_effect"])
		desc2 := parseString(row["effect"])
		if dic[id] != nil && languagesMap[lang] != "" {
			dic[id].Details[languagesMap[lang]] = desc + "\n" + desc2
		}
	})
	return result
}
