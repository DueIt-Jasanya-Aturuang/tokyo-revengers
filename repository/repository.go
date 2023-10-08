package repository

import (
	"encoding/json"
	"net/http"

	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/util"
)

type Header struct {
	Authorization string
	AppID         string
	UserID        string
	ProfileID     string
}

type QueryParam struct {
	Order     string
	Cursor    string
	StartTime string
	EndTime   string
	Type      string
}
type AllServiceRepoImpl struct {
	Account *AccountRepositoryImpl
	User    *AuthRepositoryImpl
	Finance *FinanceRepositoryImpl
}

func NewAllServiceRepoImpl() *AllServiceRepoImpl {
	return &AllServiceRepoImpl{
		Account: &AccountRepositoryImpl{},
		User:    &AuthRepositoryImpl{},
		Finance: &FinanceRepositoryImpl{},
	}
}

var ResponseGateway = map[string]any{
	"data":    nil,
	"errors":  "bad gateway",
	"message": "bad gateway",
	"status":  response.CM09,
}

func FetchResponse(r *http.Response) map[string]any {
	var resMap map[string]any
	err := json.NewDecoder(r.Body).Decode(&resMap)
	if err != nil {
		log.Warn().Msgf(util.LogErrDecode, r.Body, err)
		return ResponseGateway
	}

	return resMap
}
