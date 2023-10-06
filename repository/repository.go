package repository

import (
	"encoding/json"
	"net/http"

	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
)

type Header struct {
	Authorization string
	AppID         string
	UserID        string
	ProfileID     string
	ApiKey        string
}
type AllServiceRepoImpl struct {
	Profile *ProfileRepositoryImpl
}

func NewAllServiceRepoImpl() *AllServiceRepoImpl {
	return &AllServiceRepoImpl{
		Profile: &ProfileRepositoryImpl{},
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
		// log.Warn().Msgf(util.LogErrDecode, r.Body, err)
		return ResponseGateway
	}

	return resMap
}
