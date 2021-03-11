package main

import (
	"encoding/json"
	"flag"
	"os"

	"../story"
)

func main() {
	fileName := flag.String("json", "gopher.json", "Json file with a story")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	var story story.Story

	json.NewDecoder(file).Decode(&story)

}
