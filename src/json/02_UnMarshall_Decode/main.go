package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Fruits struct {
	Fruit  string
	Size   string
	Colour string // this will not be unmarshall since name in json file and in go different
}

func main() {
	f1, err := os.Open("example1.json")
	if err != nil {
		log.Println(err)
	}
	defer f1.Close()

	b1, err := ioutil.ReadAll(f1)
	if err != nil {
		log.Println(err)
	}
	var fruits Fruits
	json.Unmarshal(b1, &fruits)
	fmt.Println("From UnMarshal ", fruits)

	json.NewDecoder(f1).Decode(&fruits)
	fmt.Println(" From Decoder ", fruits)
}
