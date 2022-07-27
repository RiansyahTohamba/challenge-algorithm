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

func SeedBook() {
	BookList = []Book{
		{"naruto 20"},
		{"naruto 40"},
		{"naruto 22"},
		{"naruto 50"},
	}
}

// seed data book
// pakai in memory data dulu

// Create Book?
// Retrieve Book?
// Update Book

// semuanya diset di root query, apakah bisa digeser?
