package types

import (
	"github.com/graphql-go/graphql"
)

var BalanceType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BalanceType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"profile_id": &graphql.Field{
			Type: graphql.String,
		},
		"total_income_amount": &graphql.Field{
			Type: graphql.Int,
		},
		"total_income_amount_format": &graphql.Field{
			Type: graphql.String,
		},
		"total_spending_amount": &graphql.Field{
			Type: graphql.Int,
		},
		"total_spending_amount_format": &graphql.Field{
			Type: graphql.String,
		},
		"balance": &graphql.Field{
			Type: graphql.Int,
		},
		"balance_format": &graphql.Field{
			Type: graphql.String,
		},
	},
})
