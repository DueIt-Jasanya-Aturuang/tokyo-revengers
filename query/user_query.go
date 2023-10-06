package query

import (
	"net/http"

	"github.com/graphql-go/graphql"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/types"
)

func GetUser() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var user types.User
			header := p.Context.Value("headers").(http.Header)
			user.ID = header.Get("rama")
			return user, nil
		},
	}
}
