package pokeapi

type LocationArea struct {
	Count    int
	Next     string
	Previous any
	Results  []struct {
		Name string
		URL  string
	}
}

type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string
			URL  string
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int
			Version struct {
				Name string
				URL  string
			}
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int
	Location  struct {
		Name string
		URL  string
	}
	Name  string
	Names []struct {
		Language struct {
			Name string
			URL  string
		}
		Name string
	}
	PokemonEncounters []struct {
		Pokemon struct {
			Name string
			URL  string
		}
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int
				ConditionValues []any
				MaxLevel        int `json:"max_level"`
				Method          struct {
					Name string
					URL  string
				}
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string
				URL  string
			}
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
