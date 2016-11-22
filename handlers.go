package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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
	t, _ := template.ParseFiles("html/spell_list.html")
	for _, id := range arcID {
		t.Execute(w, arcDB[id])
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "The requested page does not exist - it has been lost to the Weave!")
	return
}

func SpellDisplay(w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["spellID"]
	idN, _ := strconv.Atoi(idS)
	if s, p := arcDB[idN]; p == false {
		NotFound(w, r)
	} else {
		t, _ := template.ParseFiles("html/spell.html")
		t.Execute(w, s)
	}
}
