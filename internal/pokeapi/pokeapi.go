package pokeapi

type locationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type locationArea struct {
	Id                   int        `json:"id"`
	Name                 string     `json:"name"`
	GameIndex            int        `json:"game_index"`
	EncounterMethodRates []struct{} `json:"encounter_method_rates"`
	Location             struct{}   `json:"location"`
	Names                []struct{} `json:"names"`
	PokemonEncounters    []struct {
		Pokemon        pokemon    `json:"pokemon"`
		VersionDetails []struct{} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type pokemon struct {
	ID                     int                  `json:"id"`
	Name                   string               `json:"name"`
	BaseExperience         int                  `json:"base_experience"`
	Height                 int                  `json:"height"`
	IsDefault              bool                 `json:"is_default"`
	Order                  int                  `json:"order"`
	Weight                 int                  `json:"weight"`
	Abilities              []pokemonAbility     `json:"abilities"`
	Forms                  []namedAPIResource   `json:"forms"`
	GameIndices            []versionGameIndex   `json:"game_indices"`
	HeldItems              []pokemonHeldItem    `json:"held_items"`
	LocationAreaEncounters string               `json:"location_area_encounters"`
	Moves                  []pokemonMove        `json:"moves"`
	PastTypes              []pokemonTypePast    `json:"past_types"`
	PastAbilities          []pokemonAbilityPast `json:"past_abilities"`
	Sprites                pokemonSprites       `json:"sprites"`
	Cries                  pokemonCries         `json:"cries"`
	Species                namedAPIResource     `json:"species"`
	Stats                  []pokemonStat        `json:"stats"`
	Types                  []pokemonType        `json:"types"`
}

type namedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type versionGameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   namedAPIResource `json:"version"`
}

type pokemonAbility struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  namedAPIResource `json:"ability"`
}

type pokemonAbilityPast struct {
	Generation namedAPIResource `json:"generation"`
	Abilities  []pokemonAbility `json:"abilities"`
}

type pokemonType struct {
	Slot int              `json:"slot"`
	Type namedAPIResource `json:"type"`
}

type pokemonTypePast struct {
	Generation namedAPIResource `json:"generation"`
	Types      []pokemonType    `json:"types"`
}

type pokemonHeldItem struct {
	Item          namedAPIResource         `json:"item"`
	VersionDetail []pokemonHeldItemVersion `json:"version_details"`
}

type pokemonHeldItemVersion struct {
	Version namedAPIResource `json:"version"`
	Rarity  int              `json:"rarity"`
}

type pokemonMove struct {
	Move                namedAPIResource     `json:"move"`
	VersionGroupDetails []pokemonMoveVersion `json:"version_group_details"`
}

type pokemonMoveVersion struct {
	MoveLearnMethod namedAPIResource `json:"move_learn_method"`
	VersionGroup    namedAPIResource `json:"version_group"`
	LevelLearnedAt  int              `json:"level_learned_at"`
	Order           int              `json:"order"`
}

type pokemonStat struct {
	Stat     namedAPIResource `json:"stat"`
	Effort   int              `json:"effort"`
	BaseStat int              `json:"base_stat"`
}

type pokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type pokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}
