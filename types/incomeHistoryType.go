package types

import (
	"github.com/graphql-go/graphql"
)

var IncomeHistoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "IncomeHistoryType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"profile_id": &graphql.Field{
			Type: graphql.String,
		},
		"income_type_id": &graphql.Field{
			Type: graphql.String,
		},
		"income_type_title": &graphql.Field{
			Type: graphql.String,
		},
		"payment_method_id": &graphql.Field{
			Type: graphql.String,
		},
		"payment_method_name": &graphql.Field{
			Type: graphql.String,
		},
		"income_amount": &graphql.Field{
			Type: graphql.Int,
		},
		"format_income_amount": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"time_income_history": &graphql.Field{
			Type: graphql.String,
		},
		"show_time_income_history": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var IncomeHistoryTypeList = graphql.NewObject(graphql.ObjectConfig{
	Name: "IncomeHistoryTypeList",
	Fields: graphql.Fields{
		"income_history": &graphql.Field{
			Type: graphql.NewList(IncomeHistoryType),
		},
		"cursor": &graphql.Field{
			Type: graphql.String,
		},
	},
})
