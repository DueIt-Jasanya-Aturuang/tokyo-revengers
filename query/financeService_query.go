package query

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/args"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/repository"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/types"
)

func GetPayment(finance *repository.FinanceRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.PaymentType, "PaymentResponse"),
		Args: args.InfiniteScroll,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetPayment(&repository.Header{
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

func GetSpendingType(finance *repository.FinanceRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.SpendingTypeType, "SpendingTypeListResponse"),
		Args: graphql.FieldConfigArgument{
			"order":   args.OrderInfiniteScroll,
			"cursor":  args.CursorInfiniteScroll,
			"periode": args.PeriodeBool,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			var configValue string
			if p.Args["periode"].(bool) {
				p.Args["config_name"] = "MONTHLY_PERIOD"
				resolveProfileConfig, _ := GetProfileConfig(&repository.AccountRepositoryImpl{}).Resolve(p)

				parse := resolveProfileConfig.(map[string]any)
				data, ok := parse["data"].(map[string]any)
				if !ok {
					return resolveProfileConfig, nil
				}

				configValue = data["config_value"].(string)
			}

			resp, err := finance.GetSpendingType(&repository.Header{
				Authorization: header.Get("Authorization"),
				AppID:         header.Get("App-ID"),
				UserID:        header.Get("User-ID"),
				ProfileID:     header.Get("Profile-ID"),
			}, p.Args["order"].(string), p.Args["cursor"].(string), configValue)

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

func GetDetailSpendingType(finance *repository.FinanceRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.SpendingTypeList, "SpendingTypeResponse"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetDetailSpendingType(&repository.Header{
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

func GetIncomeType(finance *repository.FinanceRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.IncomeTypeType, "IncomeTypeListResponse"),
		Args: graphql.FieldConfigArgument{
			"order":  args.OrderInfiniteScroll,
			"cursor": args.CursorInfiniteScroll,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetIncomeType(&repository.Header{
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

func GetDetailIncomeType(finance *repository.FinanceRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.IncomeListType, "IncomeTypeResponse"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetDetailIncomeType(&repository.Header{
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

func GetBalance(finance *repository.FinanceRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.BalanceType, "BalanceResponse"),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetBalance(&repository.Header{
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

func GetIncomeHistory(finance *repository.FinanceRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.IncomeHistoryTypeList, "IncomeHistoryTypeListResponse"),
		Args: graphql.FieldConfigArgument{
			"order":  args.OrderInfiniteScroll,
			"cursor": args.CursorInfiniteScroll,
			"start":  args.StartTimeParam,
			"end":    args.EndTimeParam,
			"type": &graphql.ArgumentConfig{
				Type:         args.TypeParam,
				DefaultValue: "",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetIncomeHistory(&repository.Header{
				Authorization: header.Get("Authorization"),
				AppID:         header.Get("App-ID"),
				UserID:        header.Get("User-ID"),
				ProfileID:     header.Get("Profile-ID"),
			}, &repository.QueryParam{
				Order:     p.Args["order"].(string),
				Cursor:    p.Args["cursor"].(string),
				StartTime: p.Args["start"].(string),
				EndTime:   p.Args["end"].(string),
				Type:      p.Args["type"].(string),
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

func GetDetailIncomeHistory(finance *repository.FinanceRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.IncomeHistoryType, "IncomeHistoryTypeResponse"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetDetailIncomeHistory(&repository.Header{
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

func GetSpendingHistory(finance *repository.FinanceRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.SpendingHistoryTypeList, "SpendingHistoryTypeListResponse"),
		Args: graphql.FieldConfigArgument{
			"order":  args.OrderInfiniteScroll,
			"cursor": args.CursorInfiniteScroll,
			"start":  args.StartTimeParam,
			"end":    args.EndTimeParam,
			"type": &graphql.ArgumentConfig{
				Type:         args.TypeParam,
				DefaultValue: "",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetSpendingHistory(&repository.Header{
				Authorization: header.Get("Authorization"),
				AppID:         header.Get("App-ID"),
				UserID:        header.Get("User-ID"),
				ProfileID:     header.Get("Profile-ID"),
			}, &repository.QueryParam{
				Order:     p.Args["order"].(string),
				Cursor:    p.Args["cursor"].(string),
				StartTime: p.Args["start"].(string),
				EndTime:   p.Args["end"].(string),
				Type:      p.Args["type"].(string),
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

func GetDetailSpendingHistory(finance *repository.FinanceRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.SpendingHistoryType, "SpendingHistoryTypeResponse"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetDetailSpendingHistory(&repository.Header{
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
