package query

import (
	"github.com/graphql-go/graphql"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/repository"
)

func GetRootFields(
	repo *repository.AllServiceRepoImpl,
) graphql.Fields {
	return graphql.Fields{
		"GetDetailProfile":                      GetProfile(repo.Account),
		"GetDetailUser":                         GetUser(repo.User),
		"GetDetailProfileConfig":                GetProfileConfig(repo.Account),
		"GetAllPaymentByInfiniteScroll":         GetPayment(repo.Finance),
		"GetAllSpendingTypeByInfiniteScroll":    GetSpendingType(repo.Finance),
		"GetDetailSpendingTypeByID":             GetDetailSpendingType(repo.Finance),
		"GetAllSpendingHistoryByInfiniteScroll": GetSpendingHistory(repo.Finance),
		"GetDetailSpendingHistoryByID":          GetDetailSpendingHistory(repo.Finance),
		"GetAllIncomeTypeByInfiniteScroll":      GetIncomeType(repo.Finance),
		"GetDetailIncomeTypeByID":               GetDetailIncomeType(repo.Finance),
		"GetAllIncomeHistoryByInfiniteScroll":   GetIncomeHistory(repo.Finance),
		"GetDetailIncomeHistoryByID":            GetDetailIncomeHistory(repo.Finance),
		"GetBalance":                            GetBalance(repo.Finance),
		"GetNotification":                       GetNotification(repo.Account),
		"GetDetailNotification":                 GetDetailNotification(repo.Account),
	}
}
