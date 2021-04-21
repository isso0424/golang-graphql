package hoge

import (
	"encoding/json"
	"log"

	"github.com/graphql-go/graphql"
)

type hogeArg struct {
	Value string
}

func(t hogeArg) Name() string {
	return "HogeArg"
}

func(t hogeArg) Description() string {
	return "test argument type"
}

func(t hogeArg) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return "invalid"
	}

	return string(data)
}

func(t hogeArg) Error() error {
	return nil
}

var Hoge = &graphql.Field{
	Args: graphql.FieldConfigArgument{
		"param": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Type: graphql.NewObject(
		graphql.ObjectConfig{
			Name: "hoge",
			Fields: graphql.Fields{
				"Value": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		log.Printf("%v", p.Args["param"])
		return hogeArg{ Value: p.Args["param"].(string) }, nil
	},
}
