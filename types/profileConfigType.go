package types

import (
	"github.com/graphql-go/graphql"
)

var ProfileConfigType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProfileConfigType",
	Fields: graphql.Fields{
		"profile_config_id": &graphql.Field{
			Type: graphql.String,
		},
		"profile_id": &graphql.Field{
			Type: graphql.String,
		},
		"config_name": &graphql.Field{
			Type: graphql.String,
		},
		"config_value": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"token": &graphql.Field{
			Type: graphql.String,
		},
		"days": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})
