package services

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"

	tc "github.com/kwojt/twardowski/commons"
)

const poemFilePath = "data.json"

type PoemData struct {
	Title, Body string
}

func GetRandomPoem() PoemData {
	// gather all data
	allPoems, err := ioutil.ReadFile("data.json")
	tc.Check(err)

	// serialize data
	var poems []PoemData
	err = json.Unmarshal(allPoems, &poems)
	tc.Check(err)

	// select random poem
	rand.Seed(time.Now().UnixNano())
	poem := poems[rand.Intn(len(poems))]

	return poem
}
