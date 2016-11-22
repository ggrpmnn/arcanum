package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	LIST_API_PATH  string = "/api/list/"
	SPELL_API_PATH string = "/api/spell/"
)

// a list of spell IDs, in ascending order
var arcID []int

// a list of spell structs, in random order
var arcDB map[int]Spell

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
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc(LIST_API_PATH, APIList)
	router.HandleFunc(SPELL_API_PATH+"{spellID:[0-9]+}", APISpell)
	router.HandleFunc("/spell/{spellID:[0-9]+}", SpellDisplay)
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	http.ListenAndServe(":8080", router)
}
