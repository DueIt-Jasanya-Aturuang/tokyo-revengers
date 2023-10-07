package types

import (
	"github.com/graphql-go/graphql"
)

var IncomeListType = graphql.NewObject(graphql.ObjectConfig{
	Name: "IncomeListType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"profile_id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var IncomeTypeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "IncomeTypeType",
	Fields: graphql.Fields{
		"income_type": &graphql.Field{
			Type: graphql.NewList(IncomeListType),
		},
		"cursor": &graphql.Field{Type: graphql.String},
	},
})
