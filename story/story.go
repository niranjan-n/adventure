package story

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func JsonStory(r io.Reader) (Story, error) {
	var s Story

	if err := json.NewDecoder(r).Decode(&s); err != nil {
		return nil, err
	}
	return s, nil
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chap, ok := h.s[path]; ok {
		t, _ := template.ParseFiles("story.html")
		err := t.Execute(w, chap)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "something went wrong..", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found..", http.StatusBadRequest)
}
