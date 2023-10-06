package types

import (
	"github.com/graphql-go/graphql"
)

type Profile struct {
	ProfileID string  `json:"profile_id"`
	Quote     *string `json:"quote"`
	Profesi   *string `json:"profesi"`
}

var ProfileType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Profile",

	Fields: graphql.Fields{
		"profile_id": &graphql.Field{
			Type: graphql.String,
		},
		"quote": &graphql.Field{
			Type: graphql.String,
		},
		"profesi": &graphql.Field{
			Type: graphql.String,
		},
	},
})
