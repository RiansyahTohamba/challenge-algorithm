package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func main() {
	router := gin.Default()
	router.POST("/api/graphql", func(ctx *gin.Context) {
		res, err := handleAll()
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err,
			})
		}
		ctx.JSON(200, res)
	})
	router.Run()
}

var bookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func getFields() graphql.Fields {
	books := populate()
	fields := graphql.Fields{
		"listBook": &graphql.Field{
			Type:        graphql.NewList(bookType),
			Description: "Get books list",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return books, nil
			},
		},
	}
	return fields
	// "tutorial": &graphql.Field{
	// Type: tutorialType,
	// it's good form to add a description
	// to each field.
	// Description: "Get Tutorial By ID",
	// // We can define arguments that allow us to
	// // pick specific tutorials. In this case
	// // we want to be able to specify the ID of the
	// // tutorial we want to retrieve
	// Args: graphql.FieldConfigArgument{
	// "id": &graphql.ArgumentConfig{
	// Type: graphql.Int,
	// },
	// },
	// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// // take in the ID argument
	// id, ok := p.Args["id"].(int)
	// if ok {
	// // Parse our tutorial array for the matching id
	// for _, tutorial := range tutorials {
	// if int(tutorial.ID) == id {
	// // return our tutorial
	// return tutorial, nil
	// }
	// }
	// }
	// return nil, nil
	// },
	// },
	// // this is our `list` endpoint which will return all
	// // tutorials available
	// "list": &graphql.Field{
	// Type:
	// graphql.NewList(tutorialType),
	// Description: "Get Tutorial List",
	// Resolve: func(params graphql.ResolveParams) (interface{}, error) {
	// return tutorials, nil
	// },
	// },
}
