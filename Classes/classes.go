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

var classMap = map[string]Class{}

func init() {
	classMap = make(map[string]Class)
}

func (c *Class) GetClass() *Class {
	class := &Class{}
	return class
}

// GetClasses is a function that returns a list of classes from the D&D 5e API
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

	// store the classes in a map
	for _, class := range responseObject.Classes {
		classMap[class.Index] = class
	}

	return responseObject.Classes, nil
}

func GetClassByIndex(index string) (Class, bool) {
	class, ok := classMap[index]
	return class, ok
}
