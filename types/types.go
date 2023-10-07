package types

import (
	"github.com/graphql-go/graphql"
)

func Response(output graphql.Output, name string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Fields: graphql.Fields{
			"data": &graphql.Field{
				Type: output,
			},
			"status": &graphql.Field{
				Description: "ini return string",
				Type:        graphql.String,
			},
			"message": &graphql.Field{
				Description: "ini return string",
				Type:        graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewScalar(graphql.ScalarConfig{
					Name:        "Error" + name,
					Description: "selain 400 ini return nya string, kalo 400 return map[string][]string",
					Serialize: func(value interface{}) interface{} {
						return value
					},
				}),
			},
		},
	})
}
