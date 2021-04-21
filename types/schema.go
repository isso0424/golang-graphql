package types

import (
	"isso0424/golang-graphql/types/b64enc"
	"isso0424/golang-graphql/types/hello"
	"isso0424/golang-graphql/types/hoge"
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
						"hoge": hoge.Hoge,
						"base64": b64enc.B64enc,
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
