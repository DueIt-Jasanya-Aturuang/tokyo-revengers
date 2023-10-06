package query

import (
	"net/http"

	"github.com/graphql-go/graphql"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/repository"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/types"
)

func GetProfile(profileRepo *repository.ProfileRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)
			resp, err := profileRepo.GetProfile(&repository.Header{
				Authorization: header.Get("Authorization"),
				AppID:         header.Get("App-ID"),
				UserID:        header.Get("User-ID"),
				ProfileID:     header.Get("Profile-ID"),
				ApiKey:        header.Get("X-Api-Key"),
			})

			return resp, err
		},
	}
}
