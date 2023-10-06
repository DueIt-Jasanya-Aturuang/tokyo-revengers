package query

import (
	"github.com/graphql-go/graphql"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/repository"
)

func GetRootFields(
	repo *repository.AllServiceRepoImpl,
) graphql.Fields {
	return graphql.Fields{
		"getProfile": GetProfile(repo.Profile),
		"getUser":    GetUser(repo.User),
	}
}
