package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Client struct {
    httpClient http.Client
}

func NewClient() Client {
    return Client{
        httpClient: http.Client{
            Timeout: time.Second * 10,
        },
    }
}

func (c *Client) GetLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
    endpoint := "/location-area"
    fullURL := baseURL + endpoint
    if pageUrl != nil {
        fullURL = *pageUrl
    }

    req, err := http.NewRequest(http.MethodGet, fullURL, nil)
    if err != nil {
        return LocationAreasResponse{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return LocationAreasResponse{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return LocationAreasResponse{}, err
    }

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return LocationAreasResponse{}, err
    }

    locationAreas := LocationAreasResponse{}
    err = json.Unmarshal(data, &locationAreas)
    if err != nil {
        return LocationAreasResponse{}, err
    }
    return locationAreas, nil
}
