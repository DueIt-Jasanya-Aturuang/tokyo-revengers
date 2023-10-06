package query

import (
	"net/http"

	"github.com/graphql-go/graphql"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/repository"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/types"
)

func GetProfile(profileRepo *repository.ProfileRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.ProfileType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var profile types.Profile
			header := p.Context.Value("headers").(http.Header)
			profile.ProfileID = profileRepo.Hello(header.Get("asd"))
			return profile, nil
		},
	}
}
