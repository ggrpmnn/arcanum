package main

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
type SpellEntry struct {
	Name string   `json:"name"`
	URL  string   `json:"url"`
	Tags []string `json:"tags"`
}
