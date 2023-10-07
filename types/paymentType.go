package types

import (
	"github.com/graphql-go/graphql"
)

var PaymentList = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentList",
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
		"image": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PaymentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentType",
	Fields: graphql.Fields{
		"payments": &graphql.Field{
			Type: graphql.NewList(PaymentList),
		},
		"cursor": &graphql.Field{Type: graphql.String},
	},
})
