package schema

import "github.com/graphql-go/graphql"

var JournalList []Journal

type Journal struct {
}

var journalType = graphql.NewObject(graphql.ObjectConfig{})
