package types

import (
	"github.com/graphql-go/graphql"
)

var SpendingTypeList = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingTypeList",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"profile_id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"maximum_limit": &graphql.Field{
			Type: graphql.Int,
		},
		"format_maximum_limit": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
		"used": &graphql.Field{
			Type: graphql.Int,
		},
		"format_used": &graphql.Field{
			Type: graphql.String,
		},
		"persentase_used": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var SpendingTypeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingTypeType",
	Fields: graphql.Fields{
		"spending_type": &graphql.Field{Type: graphql.NewList(SpendingTypeList)},
		"cursor":        &graphql.Field{Type: graphql.String},
		"budget_amount": &graphql.Field{
			Type: graphql.Int,
		},
		"format_budget_amount": &graphql.Field{
			Type: graphql.String,
		},
	},
})
