package types

import (
	"isso0424/golang-graphql/types/hello"
	"log"

	"github.com/graphql-go/graphql"
)

func CreateSchema() graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "RootQuery",
					Fields: graphql.Fields{
						"hello": hello.Hello,
					},
				},
			),
		},
	)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}
