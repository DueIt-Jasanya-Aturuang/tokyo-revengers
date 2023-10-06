package query

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/repository"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/types"
)

func GetProfile(profileRepo *repository.AccountRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.ProfileType, "profileResponse"),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)
			resp, err := profileRepo.GetProfile(&repository.Header{
				Authorization: header.Get("Authorization"),
				AppID:         header.Get("App-ID"),
				UserID:        header.Get("User-ID"),
				ProfileID:     header.Get("Profile-ID"),
			})

			if err != nil {
				httpResp := response.HttpResponse{
					Status:  response.CM99,
					Message: "internal server error",
					Errors:  err.Error(),
					Data:    nil,
				}
				return httpResp, nil
			}

			return resp, nil
		},
	}
}
