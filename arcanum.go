package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/template"
	//"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////
// Types
///////////////////////////////////////////////////////////////////////////////
type Spell struct {
	ID           int             `json:"id"`
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

// A list of spell IDs, in ascending order
var arcID []int

// A list of spell structs, in random order
var arcDB map[int]Spell

///////////////////////////////////////////////////////////////////////////////
// Handler functions
///////////////////////////////////////////////////////////////////////////////
func errorHandler(w http.ResponseWriter, r *http.Request, s int) {
	w.WriteHeader(s)
	if s == http.StatusNotFound {
		fmt.Fprintf(w, "Sorry - the page '%s' has been lost to the Weave!", r.URL)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	t, _ := template.ParseFiles("html/spell_list.html")
	for _, id := range arcID {
		//fmt.Fprintln(w, arcDB[id].Name)
		t.Execute(w, arcDB[id])
	}
}

func spellHandler(w http.ResponseWriter, r *http.Request) {
	idS := r.URL.Path[len("/spell/"):]
	idN, e := strconv.Atoi(idS)
	if idS == "" {
		// no ID passed
		fmt.Fprintln(w, "No id passed")
	} else if e != nil {
		// non-numeric ID passed
		fmt.Fprintln(w, "Non-numeric ID passed: ERROR")
	} else {
		// numeric ID passed; check if in DB
		if s, p := arcDB[idN]; p == false {
			fmt.Fprintln(w, "ID doesn't exist in DB!")
		} else {
			t, _ := template.ParseFiles("html/spell.html")
			t.Execute(w, s)
		}
	}
}

///////////////////////////////////////////////////////////////////////////////
// Server core functions
///////////////////////////////////////////////////////////////////////////////
func init() {
	arcID = make([]int, 0)
	arcDB = make(map[int]Spell)

	// The spell JSON files should be in the same dir as the executable
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
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/spell/", spellHandler)
	http.ListenAndServe(":8080", nil)
}

// if the error provided is not nil, panic
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
