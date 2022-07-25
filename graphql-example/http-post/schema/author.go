package schema

import "github.com/graphql-go/graphql"

var AuthorList []Author

type Author struct{}

var authorType = graphql.NewObject(graphql.ObjectConfig{})

// ?
