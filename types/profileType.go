package types

import (
	"github.com/graphql-go/graphql"
)

var ProfileType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProfileType",
	Fields: graphql.Fields{
		"profile_id": &graphql.Field{
			Type: graphql.String,
		},
		"quote": &graphql.Field{
			Type: graphql.String,
		},
		"profesi": &graphql.Field{
			Type: graphql.String,
		},
	},
})
