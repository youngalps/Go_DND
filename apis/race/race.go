package race

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Race represents a D&D race
type Race struct {
	Name  string `json:"name"`
	Index string `json:"index"`
	// Add other fields as needed
}

// GetRace fetches data for a specific D&D race by its index
func GetRace(index string) (*Race, error) {
	url := fmt.Sprintf("https://www.dnd5eapi.co/api/races/%s", index)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response into Race struct
	var race Race
	err = json.Unmarshal(body, &race)
	if err != nil {
		return nil, err
	}

	return &race, nil
}
