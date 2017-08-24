package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// APIList is the handler function for /api/list
func APIList(w http.ResponseWriter, r *http.Request) {
	l := make([]SpellEntry, 0)
	for i := 0; i < len(arcID); i++ {
		s := SpellEntry{Name: arcDB[arcID[i]].Name, URL: "http://" + r.Host + SpellEndpoint + strconv.Itoa(arcID[i]), Tags: arcDB[arcID[i]].Tags}
		l = append(l, s)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(l)
}

// APISpell is the handler function for /api/spell
func APISpell(w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["spellID"]
	// router handles non-numeric IDs; no need to error check
	idN, _ := strconv.Atoi(idS)
	if s, p := arcDB[idN]; p == false {
		notFound(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(s)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "header", nil)
	templates.ExecuteTemplate(w, "spell_list", arcDB)
	templates.ExecuteTemplate(w, "footer", nil)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "The requested page does not exist - it has been lost to the Weave!")
}

func spellDisplay(w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["spellID"]
	// router handles non-numeric IDs; no need to error check
	idN, _ := strconv.Atoi(idS)
	if s, p := arcDB[idN]; p == false {
		notFound(w, r)
	} else {
		templates.ExecuteTemplate(w, "header", nil)
		templates.ExecuteTemplate(w, "spell", s)
		// convert newline chars to <br> for proper HTML display
		d := template.HTML(strings.Replace(s.Description, "\n", "<br>", -1))
		templates.ExecuteTemplate(w, "spell_desc", d)
		templates.ExecuteTemplate(w, "footer", nil)
	}
}
