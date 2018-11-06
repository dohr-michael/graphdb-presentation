package pokeapi

type Id int

type Labels map[string]string

type WithIdentity interface {
	GetId() Id
	SetName(lang string, trans string)
}

type Identity struct {
	Id         Id     `json:"id"`
	Identifier string `json:"identifier"`
	Name       Labels `json:"name"`
}

func (t *Identity) GetId() Id {
	return t.Id
}

func (t *Identity) SetName(lang string, trans string) {
	t.Name[lang] = trans
}

func NewIdentity(id int, identifier string) Identity {
	return Identity{
		Id:         Id(id),
		Identifier: identifier,
		Name:       make(Labels),
	}
}

// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/evolution_triggers.csv
// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/evolution_trigger_prose.csv
type EvolutionTrigger struct {
	Identity
}

// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon_evolution.csv
type Evolution struct {
	Into         Id  `json:"into"`
	TriggerType  Id  `json:"triggerType"`
	TriggerItem  Id  `json:"triggerItem,omitempty"`
	MinimumLevel int `json:"minimumLevel,omitempty"`
}

// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/generations.csv
// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/generation_names.csv
type Generation struct {
	Identity
}

// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/type_names.csv
// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/types.csv
type Type struct {
	Identity
}

// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon.csv
// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon_types.csv
// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon_species.csv
// https://github.com/PokeAPI/pokeapi/blob/master/data/v2/csv/pokemon_species_names.csv
type Pokemon struct {
	Identity
	Weight     int         `json:"weight"`
	Height     int         `json:"height"`
	BaseExp    int         `json:"baseExp"`
	Generation Id          `json:"generation"`
	Types      []Id        `json:"types"`
	FightWith  []MoveLink  `json:"-"`
	EvolveInto []Evolution `json:"evolveInto"`
}

// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon_moves.csv
type MoveLink struct {
	MoveId     Id  `json:"moveId"`
	MoveMethod Id  `json:"moveMethod"`
	Level      int `json:"level"`
}

// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon_move_methods.csv
// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/pokemon_move_method_prose.csv
type MoveMethod struct {
	Identity
}

// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/moves.csv
// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/move_names.csv
type Move struct {
	Identity
	Generation Id  `json:"generation"`
	Type       Id  `json:"type"`
	Power      int `json:"power"`
	Pp         int `json:"pp"`
	Accuracy   int `json:"accuracy"`
}

// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/item_pockets.csv
// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/item_pocket_names.csv
type Pocket struct {
	Identity
}

// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/item_categories.csv
// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/item_category_prose.csv
type ItemCategory struct {
	Identity
	PocketId Id `json:"pocketId"`
}

// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/items.csv
// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/item_names.csv
// https://github.com/PokeAPI/pokeapi/raw/master/data/v2/csv/item_prose.csv
type Item struct {
	Identity
	Details  Labels `json:"details"`
	Category Id     `json:"category"`
	Cost     int    `json:"cost"`
}
