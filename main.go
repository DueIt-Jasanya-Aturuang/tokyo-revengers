package main

import (
	"context"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"

	"github.com/graphql-go/handler"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/infra"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/query"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/repository"
)

func main() {
	infra.InitEnv()
	repo := repository.NewAllServiceRepoImpl()
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: query.GetRootFields(repo),
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed new scheman graph | err %v", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema: &schema,
		FormatErrorFn: func(err error) gqlerrors.FormattedError {
			return gqlerrors.FormattedError{
				Message:    "",
				Locations:  nil,
				Path:       nil,
				Extensions: nil,
			}
		},
	})

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "headers", r.Header)
		r = r.WithContext(ctx)
		w.Header().Set("Authorization", r.Header.Get("Authorization"))
		httpHandler.ServeHTTP(w, r)
	}))
	log.Print("ready: listening...\n")

	http.ListenAndServe(":8383", nil)
}
