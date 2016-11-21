package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// used for listing names-only
type SpellName struct {
	Name string `json:"name"`
}

// A list of spell IDs, in ascending order
var arcID []int

// A list of spell structs, in random order
var arcDB map[int]Spell

///////////////////////////////////////////////////////////////////////////////
// Handler functions
///////////////////////////////////////////////////////////////////////////////
func apiList(w http.ResponseWriter, r *http.Request) {
	l := make([]SpellName, 0)
	for i := 0; i < len(arcID); i++ {
		n := SpellName{Name: arcDB[arcID[i]].Name}
		l = append(l, n)
	}
	json.NewEncoder(w).Encode(l)
}

func apiSpell(w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["spellID"]
	idN, e := strconv.Atoi(idS)
	if e != nil {
		fmt.Fprintln(w, "NOPE")
	} else {
		s := arcDB[idN]
		json.NewEncoder(w).Encode(s)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/spell_list.html")
	for _, id := range arcID {
		t.Execute(w, arcDB[id])
	}
}

func spellDisplay(w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["spellID"]
	idN, e := strconv.Atoi(idS)
	if e != nil {
		// non-numeric ID passed; invoke 404
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
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", index)
	router.HandleFunc("/spell/{spellID}", spellDisplay)
	router.HandleFunc("/api/list", apiList)
	router.HandleFunc("/api/spell/{spellID}", apiSpell)

	http.ListenAndServe(":8080", router)
}

// if the error provided is not nil, panic
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
