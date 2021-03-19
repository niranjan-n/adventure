package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"./story"
)

func main() {
	port := flag.Int("port", 3000, "Port to start the application")
	fileName := flag.String("json", "gopher.json", "Json file with a story")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	s, err := story.JsonStory(file)
	if err != nil {
		panic(err)
	}
	h := story.NewHandler(s)
	fmt.Printf("Server started on port : %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
