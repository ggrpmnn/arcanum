package main

// Spell structs contain the full data for a spell object
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

// SpellComponents contain the data for the components of a particular spell
// for the unintiated: verbal components are spoken words or sounds, somatic
// components are movements of the hands or body, and material components are
// small items used to cast a spell; a spell may require none, any, or all of
// these pieces
type SpellComponents struct {
	Verbal         bool     `json:"verbal"`
	Somatic        bool     `json:"somatic"`
	Material       bool     `json:"material"`
	MaterialString []string `json:"materials_needed,omitempty"`
}

// SpellEntry structs are used for the list API and contain a much smaller
// subset of data than the full spell object
type SpellEntry struct {
	Name string   `json:"name"`
	URL  string   `json:"url"`
	Tags []string `json:"tags"`
}
