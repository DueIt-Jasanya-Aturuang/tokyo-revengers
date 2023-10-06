package query

import (
	"github.com/graphql-go/graphql"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/repository"
)

// GetRootFields returns all the available queries.
func GetRootFields(
	profile *repository.AllServiceRepoImpl,
) graphql.Fields {
	return graphql.Fields{
		"profile": GetProfile(profile.Profile),
		"user":    GetUser(),
	}
}
