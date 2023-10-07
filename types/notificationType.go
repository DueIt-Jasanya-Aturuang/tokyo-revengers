package types

import (
	"github.com/graphql-go/graphql"
)

var NotificationTypeList = graphql.NewObject(graphql.ObjectConfig{
	Name: "NotificationTypeList",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"profile_id": &graphql.Field{
			Type: graphql.String,
		},
		"user_config_id": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var NotificationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NotificationType",
	Fields: graphql.Fields{
		"notification": &graphql.Field{Type: graphql.NewList(NotificationTypeList)},
		"cursor":       &graphql.Field{Type: graphql.String},
	},
})
