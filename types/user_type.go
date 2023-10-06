package types

import (
	"github.com/graphql-go/graphql"
)

type User struct {
	ID          string  `json:"id"`
	FullName    string  `json:"full_name"`
	Gender      string  `json:"gender"`
	Image       string  `json:"image"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	EmailFormat string  `json:"email_format"`
	PhoneNumber *string `json:"phone_number"`
	Activited   bool    `json:"activited"`
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
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
