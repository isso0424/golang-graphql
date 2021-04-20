package server

import (
	"encoding/json"
	"io/ioutil"
	"isso0424/golang-graphql/types"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

var schema = types.CreateSchema()

func graphqlRoute(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)

		return
	}
	result := graphql.Do(
		graphql.Params{
			Schema: schema,
			RequestString: string(data),
		},
	)
	if len(result.Errors) > 0 {
		log.Printf("failed to execute graphql operation, errors: %+v", result.Errors)
	}
	rJSON, _ := json.Marshal(result)
	w.Write(rJSON)
}
