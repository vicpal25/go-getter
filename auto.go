package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Brand ...
type Brand struct {
	Name string
}

//Auto ...
type Auto struct {
	Brand Brand  `json:"brand,omitempty"`
	Model string `json:"model,omitempty"`
	Price int    `json:"price,omitempty"`
}

//AutosDictionary ...
type AutosDictionary []struct {
	Brand  string   `json:"brand"`
	Models []string `json:"models"`
}

//Make ...
func Make() (Auto, error) {

	var v Auto

	Generate(v)

	return v, nil

}

//Generate ...
func Generate(c Auto) {

	var br Brand

	br.Name = "Honda"

	Collect()

	fmt.Println("Hello, playground", c)

}

func SearchBrand(c Auto) Brand {

	jsonFile, err := os.Open("cars.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var autos AutosDictionary

	json.Unmarshal(byteValue, &autos)

	for i := 0; i < len(autos); i++ {
		fmt.Println("User Type: " + autos[i].Brand)

	}

}
