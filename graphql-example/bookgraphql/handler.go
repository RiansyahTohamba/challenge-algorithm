package main

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"
)

func handleAll(query string) (*graphql.Result, error) {
	fields := getFields()

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	scConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(scConfig)

	if err != nil {
		log.Fatalf("failed create schema: %v", err)
	}

	// query := `
	// 	{
	// 		listBook{
	// 			title
	// 			author
	// 		}
	// 	}
	// `
	params := graphql.Params{Schema: schema, RequestString: query}

	res := graphql.Do(params)
	if len(res.Errors) > 0 {
		log.Fatalf("failed execute graphql operation %v", res.Errors)
		return nil, errors.New("there is problem")
	}

	return res, nil
}
