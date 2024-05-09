package classes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Class struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func GetClasses() ([]Class, error) {
	response, err := http.Get("https://www.dnd5eapi.co/api/classes")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject struct {
		Classes []Class `json:"results"`
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}

	return responseObject.Classes, nil
}
