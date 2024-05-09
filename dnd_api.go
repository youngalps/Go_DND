package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Classes struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func main() {
	response, err := http.Get("https://www.dnd5eapi.co/api/classes")

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var classes []Classes
	err = json.Unmarshal(responseData, &classes)
	if err != nil {
		log.Fatal(err)
	}

	for _, class := range classes {
		fmt.Println(class.Index)
		fmt.Println(class.Name)
		fmt.Println(class.URL)
		fmt.Println()
	}

}
