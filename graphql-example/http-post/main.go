package main

import (
	"challenge-algorithm/graphql-example/http-post/schema"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

// contoh reques di postman
func main() {
	// seedData book
	schema.SeedBook()
	schema.SeedJournal()
	http.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		var post postData

		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			w.WriteHeader(400)
			return
		}
		// bagaimana kalau buat all schema?
		result := graphql.Do(graphql.Params{
			Context:        r.Context(),
			Schema:         schema.MainSchema,
			RequestString:  post.Query,
			VariableValues: post.Variables,
			OperationName:  post.Operation,
		})

		if err := json.NewEncoder(w).Encode(result); err != nil {
			fmt.Printf("could not write result to response: %s", err)
		}
	})

	fmt.Println("running on port 8080")
	http.ListenAndServe(":8080", nil)
}
