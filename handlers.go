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
	json.NewEncoder(w).Encode(l)
}

func APISpell(w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["spellID"]
	idN, e := strconv.Atoi(idS)
	if e != nil {
		fmt.Fprintln(w, "NOPE")
	} else {
		s := arcDB[idN]
		json.NewEncoder(w).Encode(s)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/spell_list.html")
	for _, id := range arcID {
		t.Execute(w, arcDB[id])
	}
}

func SpellDisplay(w http.ResponseWriter, r *http.Request) {
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
