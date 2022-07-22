package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	_ = importJSONDataFromFile("data.json", &data)
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {

		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	// bisa via get juga
}

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	res := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(res.Errors) > 0 {
		fmt.Printf("wrong result, un-expected errors: %v", res.Errors)
	}
	return res
}

func importJSONDataFromFile(filename string, result interface{}) (isOK bool) {
	isOK = true
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print("Error:", err)
		isOK = false

	}
	err = json.Unmarshal(contents, result)
	if err != nil {
		fmt.Print("error:", err)
		isOK = false
	}
	return
}
