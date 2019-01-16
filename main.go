package main

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Start Return Types for User
var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"accountList": &graphql.Field{
			Type: graphql.NewList(accountType),
		},
	},
})

var accountType = graphql.NewObject(graphql.ObjectConfig{
	Name: "accountList",
	Fields: graphql.Fields{
		"accNo": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"balance": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Float),
		},
	},
})
// End User Type

// Address Type
var addressType = graphql.NewObject(graphql.ObjectConfig{
	Name: "address",
	Fields: graphql.Fields{
		"Num": &graphql.Field{
			Type: graphql.String,
		},
		"Street": &graphql.Field{
			Type: graphql.String,
		},
		"Type": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"lastestPost": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Hello Desc",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello World!", nil
			},
		},
		"random": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "Returns a random number",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return r.Int31(), nil
			},
		},
		"double": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "Doubles the input number",
			Args: graphql.FieldConfigArgument{
				"val": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				val, _ := p.Args["val"].(int)

				return val * 2, nil
			},
		},
		"user": &graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// Read JSON, and spit out GraphQL
				response, err := http.Get("http://localhost:4545/test")
				if err != nil {
					return nil, err
				}

				var dat map[string]interface{}
				data, _ := ioutil.ReadAll(response.Body)

				if err := json.Unmarshal(data, &dat); err != nil {
					return nil, err
				}

				return dat, nil
			},
		},
		"address": &graphql.Field{
			Type: addressType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// Read JSON, and spit out GraphQL
				response, err := http.Get("http://localhost:4546/address")
				if err != nil {
					return nil, err
				}

				var dat map[string]interface{}
				data, _ := ioutil.ReadAll(response.Body)

				if err := json.Unmarshal(data, &dat); err != nil {
					return nil, err
				}

				return dat["Address"], nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func main() {
	h := handler.New(&handler.Config{
		Schema:     &Schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", h)

	http.ListenAndServe(":8090", nil)
}
