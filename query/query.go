package query

import (
	"github.com/graphql-go/graphql"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/repository"
)

func GetRootFields(
	repo *repository.AllServiceRepoImpl,
) graphql.Fields {
	return graphql.Fields{
		"GetDetailProfile":                   GetProfile(repo.Profile),
		"GetDetailUser":                      GetUser(repo.User),
		"GetDetailProfileConfig":             GetProfileConfig(repo.Profile),
		"GetAllPaymentByInfiniteScroll":      GetPayment(repo.Finance),
		"GetAllSpendingTypeByInfiniteScroll": GetSpendingType(repo.Finance),
		"GetDetailSpendingTypeByID":          GetDetailSpendingType(repo.Finance),
		"GetAllIncomeTypeByInfiniteScroll":   GetIncomeType(repo.Finance),
		"GetDetailIncomeTypeByID":            GetDetailIncomeType(repo.Finance),
	}
}
