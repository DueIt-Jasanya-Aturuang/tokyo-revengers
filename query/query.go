package query

import (
	"github.com/graphql-go/graphql"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/repository"
)

func GetRootFields(
	profile *repository.AllServiceRepoImpl,
) graphql.Fields {
	return graphql.Fields{
		"getProfile": GetProfile(profile.Profile),
		"user":       GetUser(),
	}
}
