package types

import (
	"github.com/graphql-go/graphql"
)

var SpendingHistoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingHistoryType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"profile_id": &graphql.Field{
			Type: graphql.String,
		},
		"spending_type_id": &graphql.Field{
			Type: graphql.String,
		},
		"spending_type_title": &graphql.Field{
			Type: graphql.String,
		},
		"payment_method_id": &graphql.Field{
			Type: graphql.String,
		},
		"payment_method_name": &graphql.Field{
			Type: graphql.String,
		},
		"payment_name": &graphql.Field{
			Type: graphql.String,
		},
		"before_balance": &graphql.Field{
			Type: graphql.Int,
		},
		"spending_amount": &graphql.Field{
			Type: graphql.Int,
		},
		"format_spending_amount": &graphql.Field{
			Type: graphql.String,
		},
		"after_balance": &graphql.Field{
			Type: graphql.Int,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"time_spending_history": &graphql.Field{
			Type: graphql.String,
		},
		"show_time_spending_history": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var SpendingHistoryTypeList = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingHistoryTypeList",
	Fields: graphql.Fields{
		"spending_history": &graphql.Field{
			Type: graphql.NewList(SpendingHistoryType),
		},
		"cursor": &graphql.Field{
			Type: graphql.String,
		},
	},
})
