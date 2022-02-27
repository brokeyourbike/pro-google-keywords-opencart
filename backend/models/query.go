package models

import "time"

type Query struct {
	LanguageCode string
	QueryString  string
	CreateDate   time.Time
	UpdateDate   time.Time
	Suggestions  map[string]Suggestion
}

type Suggestion struct {
	Type      string
	Relevance int
}
