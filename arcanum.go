package main

import (
	"fmt"
	"io/ioutil"
	//"html/template"
	"encoding/json"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////
// Types
///////////////////////////////////////////////////////////////////////////////
type Spell struct {
	Name         string          `json:"name"`
	Tags         []string        `json:"tags"`
	Type         string          `json:"type"`
	Time         string          `json:"casting_time"`
	Range        string          `json:"range"`
	Components   SpellComponents `json:"components"`
	Duration     string          `json:"duration"`
	Description  string          `json:"description"`
	HigherLevels string          `json:"higher_levels,omitempty"`
}

type SpellComponents struct {
	Verbal         bool     `json:"verbal"`
	Somatic        bool     `json:"somatic"`
	Material       bool     `json:"material"`
	MaterialString []string `json:"materials_needed,omitempty"`
}

var arcDB []Spell

///////////////////////////////////////////////////////////////////////////////
// Handler functions
///////////////////////////////////////////////////////////////////////////////
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Welcome to Arcanum! Spells go here!")
	for i := range arcDB {
		fmt.Fprintln(w, arcDB[i].Name)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, s int) {
	w.WriteHeader(s)
	if s == http.StatusNotFound {
		fmt.Fprintf(w, "Sorry - the page '%s' has been lost to the Weave!", r.URL)
	}
}

///////////////////////////////////////////////////////////////////////////////
// Server core functions
///////////////////////////////////////////////////////////////////////////////
func init() {
	arcDB = make([]Spell, 0)

	// The spell JSON files should be in the same dir as the executable
	//spellFiles := []string{"phb_spells.json", "eepc_spells.json", "scag_spells.json"}
	spellFiles := []string{"test_spells.json"}
	for _, f := range spellFiles {
		j, err := ioutil.ReadFile(f)
		checkError(err)
		sl := make([]Spell, 0)
		err = json.Unmarshal(j, &sl)
		checkError(err)
		// Copy the new list into the DB list
		for i := 0; i < len(sl); i++ {
			arcDB = append(arcDB, sl[i])
		}
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

// if the error provided is not nil, panic
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
