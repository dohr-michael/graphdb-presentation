package importer


import (
	"github.com/dohr-michael/graphdb-presentation/pokeapi"
)

func LoadEvolutionTriggers() interface{} {
	var result []*pokeapi.EvolutionTrigger
	translationLoader(loaderParams{
		BaseUrl:              "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/evolution_triggers.csv",
		TranslationUrl:       "https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/evolution_trigger_prose.csv",
		IdentifierField:      "identifier",
		IdField:              "id",
		TranslationIdField:   "evolution_trigger_id",
		TranslationLangField: "local_language_id",
		TranslationField:     "name",
		Factory: func(id int, identifier string) pokeapi.WithIdentity {
			return &pokeapi.EvolutionTrigger{
				Identity: pokeapi.NewIdentity(id, identifier),
			}
		},
	}, func(item pokeapi.WithIdentity, row map[string]string) {
		result = append(result, item.(*pokeapi.EvolutionTrigger))
	})

	return result
}
