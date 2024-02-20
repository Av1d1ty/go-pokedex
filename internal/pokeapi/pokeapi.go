package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Av1d1ty/go-pokedex/internal/pokekache"
)

const baseURL = "https://pokeapi.co/api/v2"

type LocationAreaListResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type Client struct {
    httpClient http.Client
    cache pokekache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
    return Client{
        cache: pokekache.NewCache(cacheInterval),
        httpClient: http.Client{
            Timeout: time.Second * 10,
        },
    }
}

func (c *Client) GetLocationAreaList(pageUrl *string) (LocationAreaListResponse, error) {
    endpoint := "/location-area"
    fullURL := baseURL + endpoint
    if pageUrl != nil {
        fullURL = *pageUrl
    }

    if data, ok := c.cache.Get(fullURL); ok {
        locationAreas := LocationAreaListResponse{}
        err := json.Unmarshal(data, &locationAreas)
        if err != nil {
            return LocationAreaListResponse{}, err
        }
        return locationAreas, nil
    }

    req, err := http.NewRequest(http.MethodGet, fullURL, nil)
    if err != nil {
        return LocationAreaListResponse{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return LocationAreaListResponse{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return LocationAreaListResponse{}, err
    }

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return LocationAreaListResponse{}, err
    }

    locationAreas := LocationAreaListResponse{}
    err = json.Unmarshal(data, &locationAreas)
    if err != nil {
        return LocationAreaListResponse{}, err
    }

    c.cache.Set(fullURL, data)

    return locationAreas, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
    endpoint := "/location-area/" + locationAreaName
    fullURL := baseURL + endpoint

    if data, ok := c.cache.Get(fullURL); ok {
        locationAreas := LocationArea{}
        err := json.Unmarshal(data, &locationAreas)
        if err != nil {
            return LocationArea{}, err
        }
        return locationAreas, nil
    }

    req, err := http.NewRequest(http.MethodGet, fullURL, nil)
    if err != nil {
        return LocationArea{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return LocationArea{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return LocationArea{}, err
    }

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return LocationArea{}, err
    }

    locationAreas := LocationArea{}
    err = json.Unmarshal(data, &locationAreas)
    if err != nil {
        return LocationArea{}, err
    }

    c.cache.Set(fullURL, data)

    return locationAreas, nil
}
