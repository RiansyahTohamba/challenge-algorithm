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
		// mutation+_{createBook(title:"Naruto+20"){id,title}}
		"createBook": &graphql.Field{
			Type:        bookType,
			Description: "Create New Book",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		},
		"createAuthor": &graphql.Field{
			Type:        authorType,
			Description: "Create New Author",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		},
		"createJournal": &graphql.Field{
			Type:        journalType,
			Description: "Create New Journal",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
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
			Type:        graphql.NewList(journalType),
			Description: "List of Journal",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return JournalList, nil
			},
		},
	},
})

var MainSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
