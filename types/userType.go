package types

import (
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserType",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.Int},
		"full_name":    &graphql.Field{Type: graphql.String},
		"gender":       &graphql.Field{Type: graphql.String},
		"image":        &graphql.Field{Type: graphql.String},
		"username":     &graphql.Field{Type: graphql.String},
		"email":        &graphql.Field{Type: graphql.String},
		"email_format": &graphql.Field{Type: graphql.String},
		"phone_number": &graphql.Field{Type: graphql.String},
		"activited":    &graphql.Field{Type: graphql.Boolean},
	},
})
