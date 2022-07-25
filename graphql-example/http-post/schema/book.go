package schema

import "github.com/graphql-go/graphql"

var BookList []Book

type Book struct {
	Title string `json:"title"`
}

var bookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "book",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
	},
})
