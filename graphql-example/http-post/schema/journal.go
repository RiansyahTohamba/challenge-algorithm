package schema

import "github.com/graphql-go/graphql"

var JournalList []Journal

type Journal struct {
	Title string `json:"title"`
}

var journalType = graphql.NewObject(graphql.ObjectConfig{
	Name: "journal",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func SeedJournal() {
	JournalList = []Journal{
		{""},
		{""},
	}
}
