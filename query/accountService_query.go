package query

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/args"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/repository"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/types"
)

func GetProfile(profileRepo *repository.AccountRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.ProfileType, "ProfileResponse"),
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

func GetProfileConfig(account *repository.AccountRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.ProfileConfigType, "ProfileConfigResponse"),
		Args: graphql.FieldConfigArgument{
			"config_name": &graphql.ArgumentConfig{
				Type:         args.ProfileConfigEnum,
				DefaultValue: "MONTHLY_PERIOD",
				Description:  "Input konfigurasi profil Config, by default MONTHLY_PERIOD",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := account.GetProfileConfig(&repository.Header{
				Authorization: header.Get("Authorization"),
				AppID:         header.Get("App-ID"),
				UserID:        header.Get("User-ID"),
				ProfileID:     header.Get("Profile-ID"),
			}, p.Args["config_name"].(string))

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

func GetNotification(account *repository.AccountRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.NotificationType, "NotificationResponse"),
		Args: graphql.FieldConfigArgument{
			"order":  args.OrderInfiniteScroll,
			"cursor": args.CursorInfiniteScroll,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := account.GetNotification(&repository.Header{
				Authorization: header.Get("Authorization"),
				AppID:         header.Get("App-ID"),
				UserID:        header.Get("User-ID"),
				ProfileID:     header.Get("Profile-ID"),
			}, p.Args["order"].(string), p.Args["cursor"].(string))

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

func GetDetailNotification(account *repository.AccountRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.NotificationTypeList, "NotificationListResponse"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := account.GetDetailNotification(&repository.Header{
				Authorization: header.Get("Authorization"),
				AppID:         header.Get("App-ID"),
				UserID:        header.Get("User-ID"),
				ProfileID:     header.Get("Profile-ID"),
			}, p.Args["id"].(string))

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
