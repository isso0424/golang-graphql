package b64enc

import (
	"encoding/base64"
	"errors"

	"github.com/graphql-go/graphql"
)

var B64enc = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"base": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"isEncode": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
			DefaultValue: false,
		},
		"isDecode": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
			DefaultValue: false,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		isEncode := p.Args["isEncode"].(bool)
		isDecode := p.Args["isDecode"].(bool)
		if (!isEncode && !isDecode) {
			return nil, errors.New("isEncode or isDecode must be true.")
		}

		if (isEncode && isDecode) {
			return nil, errors.New("Cannot assign true for both isEncode and isDecode.")
		}

		if (isEncode) {
			data := p.Args["base"].(string)
			return base64.RawStdEncoding.EncodeToString([]byte(data)), nil
		}

		data := p.Args["base"].(string)
		dist, err := base64.RawStdEncoding.DecodeString(data)
		if err != nil {
			return nil, err
		}
		return string(dist), nil
	},
}
