package schema

import "github.com/graphql-go/graphql"

var AuthorList []Author

type Author struct {
	Name string `json:"name"`
}

var authorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "author",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})
