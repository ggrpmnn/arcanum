package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func APIList(w http.ResponseWriter, r *http.Request) {
	l := make([]SpellEntry, 0)
	for i := 0; i < len(arcID); i++ {
		n := arcDB[arcID[i]].Name
		u := "http://" + r.Host + SPELL_API_PATH + strconv.Itoa(arcID[i])
		t := arcDB[arcID[i]].Tags
		s := SpellEntry{Name: n, URL: u, Tags: t}
		l = append(l, s)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(l)
}

func APISpell(w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["spellID"]
	// router handles non-numeric IDs; no need to error check
	idN, _ := strconv.Atoi(idS)
	if s, p := arcDB[idN]; p == false {
		NotFound(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(s)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "header", nil)
	for _, id := range arcID {
		templates.ExecuteTemplate(w, "spell_list", arcDB[id])
	}
	templates.ExecuteTemplate(w, "footer", nil)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "The requested page does not exist - it has been lost to the Weave!")
}

func SpellDisplay(w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["spellID"]
	// router handles non-numeric IDs; no need to error check
	idN, _ := strconv.Atoi(idS)
	if s, p := arcDB[idN]; p == false {
		NotFound(w, r)
	} else {
		templates.ExecuteTemplate(w, "header", nil)
		templates.ExecuteTemplate(w, "spell", s)
		templates.ExecuteTemplate(w, "footer", nil)
	}
}
