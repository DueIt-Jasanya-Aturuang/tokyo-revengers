package types

import (
	"github.com/graphql-go/graphql"
)

var Response = graphql.NewObject(graphql.ObjectConfig{
	Name: "response",

	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: graphql.NewScalar(graphql.ScalarConfig{
				Name:        "DataResult",
				Description: "return nya sesuai endpoint yang kalian pakai",
				Serialize: func(value interface{}) interface{} {
					return value
				},
			}),
		},
		"status": &graphql.Field{
			Name:        "Status",
			Description: "ini return string",
			Type:        graphql.String,
		},
		"message": &graphql.Field{
			Name:        "Message",
			Description: "ini return string",
			Type:        graphql.String,
		},
		"errors": &graphql.Field{
			Type: graphql.NewScalar(graphql.ScalarConfig{
				Name:        "ErrorResponse", // Atur nama tipe data skalar sesuai kebutuhan Anda
				Description: "selain 400 ini return nya string, kalo 400 return map[string][]string",
				Serialize: func(value interface{}) interface{} {
					return value
				},
			}),
		},
	},
})
