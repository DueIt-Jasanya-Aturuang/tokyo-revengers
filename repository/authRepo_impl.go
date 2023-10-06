package repository

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/infra"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/util"
)

type AuthRepositoryImpl struct {
}

func (u *AuthRepositoryImpl) GetUser(header *Header) (map[string]any, error) {
	route := fmt.Sprintf("%s/auth/user", infra.AuthUrl)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("User-ID", header.UserID)
	request.Header.Set("App-ID", header.AppID)
	request.Header.Set("Profile-ID", header.ProfileID)
	request.Header.Set("Authorization", header.Authorization)
	request.Header.Set("X-Api-Key", infra.AuthKey)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Warn().Msgf(util.LogErrClientDo, err)
		return nil, err
	}
	defer func() {
		if errBody := response.Body.Close(); errBody != nil {
			log.Warn().Msgf(util.LogErrClientDoClose, err)
		}
	}()

	resp := FetchResponse(response)

	return resp, nil
}
