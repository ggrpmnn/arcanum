package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// SpellListEndpoint is the route for the spell listing
	SpellListEndpoint string = "/api/list/"
	// SpellEndpoint is the route for an individual spell
	SpellEndpoint string = "/api/spell/"
)

var (
	// arcID is a list of spell IDs, in ascending order
	arcID []int
	// arcDB is a list of spell structs (randomized order)
	arcDB map[int]Spell
	// HTML templates for rendering pages
	templates *template.Template
)

// if the error provided is not nil, panic
// for use in the init funtion only; don't want a runtime panic
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	arcID = make([]int, 0)
	arcDB = make(map[int]Spell)

	// Parse source data
	// the paths can be anywhere, but need to be relative to the source dir
	spellFiles := []string{"./data/phb_spells.json", "./data/eepc_spells.json", "./data/scag_spells.json"}
	for _, f := range spellFiles {
		j, err := ioutil.ReadFile(f)
		checkError(err)
		sl := make([]Spell, 0)
		err = json.Unmarshal(j, &sl)
		checkError(err)
		// Copy the new list into the DB list
		for i := 0; i < len(sl); i++ {
			arcID = append(arcID, sl[i].ID)
			arcDB[sl[i].ID] = sl[i]
		}
	}
	log.Printf("> Spell data created.\n")

	// Parse HTML templates
	templates = template.Must(template.ParseGlob("html/*"))
	log.Printf("> HTML templates parsed.\n")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", index)
	router.HandleFunc(SpellListEndpoint, APIList)
	router.HandleFunc(SpellEndpoint+"{spellID:[0-9]+}", APISpell)
	router.HandleFunc("/spell/{spellID:[0-9]+}", spellDisplay)
	router.NotFoundHandler = http.HandlerFunc(notFound)

	// for static files (CSS etc.)
	router.Handle("/css/{file}", http.FileServer(http.Dir("")))

	log.Printf("> Listening on port 8080.\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
