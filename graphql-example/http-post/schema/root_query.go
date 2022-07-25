package schema

import "github.com/graphql-go/graphql"

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"bookList": &graphql.Field{
			Type:        graphql.NewList(bookType),
			Description: "List of Books",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return BookList, nil
			},
		},
		"authorList": &graphql.Field{
			Type:        graphql.NewList(authorType),
			Description: "List of author",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return AuthorList, nil
			},
		},
		"journalList": &graphql.Field{
			Type: graphql.NewList(journalType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return JournalList, nil
			},
		},
	},
})

var MainSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
