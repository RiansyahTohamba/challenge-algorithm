package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	// router := gin.Default()
	// print the data
	// show via gin
	fields := graphql.Fields{
		"books": &graphql.Field{},
	}

	// RootQuery adalah syntax
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error : %v", err)
	}

	query := `
		{ 
			books 
		}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	res := graphql.Do(params)

	if len(res.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, error : %v", err)
	}

	rJSON, _ := json.Marshal(res)
	fmt.Printf("%s \n", rJSON)
}
