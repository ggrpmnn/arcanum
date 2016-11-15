package main

import (
	"fmt"
	//"html/template"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////
// Types
///////////////////////////////////////////////////////////////////////////////
type Spell struct {
	Name       string
	Tags       []string
	Type       string
	Time       string
	Range      string
	Components SpellComponents
}

type SpellComponents struct {
	Verbal         bool
	Somatic        bool
	Material       bool
	MaterialString []string
}

///////////////////////////////////////////////////////////////////////////////
// Handler functions
///////////////////////////////////////////////////////////////////////////////
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Welcome to Arcanum! Spells go here!")
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

}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
