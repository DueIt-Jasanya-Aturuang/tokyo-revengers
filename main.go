package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/infra"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/query"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/repository"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/util"
)

func main() {
	infra.InitEnv()
	infra.LogInit()

	repo := repository.NewAllServiceRepoImpl()
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: query.GetRootFields(repo),
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		// log.Fatalf("failed new scheman graph | err %v", err)
		log.Fatal().Msgf("failed new scheman graph | err %v", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema: &schema,
	})

	http.Handle("/page", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Api-Key") != infra.Key {
			httpResp := response.HttpResponse{
				Status:  response.CM05,
				Message: "forbidden",
				Errors:  "invalid api key page",
				Data:    nil,
			}

			w.WriteHeader(403)
			if errEncode := json.NewEncoder(w).Encode(httpResp); errEncode != nil {
				log.Err(errEncode).Msgf(util.LogErrEncode, httpResp, errEncode)
			}
			return
		}

		ctx := context.WithValue(r.Context(), "headers", r.Header)
		r = r.WithContext(ctx)

		w.Header().Set("Authorization", r.Header.Get("Authorization"))

		httpHandler.ServeHTTP(w, r)
	}))

	http.ListenAndServe(":8383", nil)
}
