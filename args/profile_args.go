package args

import (
	"github.com/graphql-go/graphql"
)

var ProfileArgs = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProfileArgs",
	Fields: graphql.Fields{
		"profile_args": &graphql.Field{
			Type: graphql.String,
		},
	},
})
