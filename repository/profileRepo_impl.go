package repository

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/infra"
)

type ProfileRepositoryImpl struct {
}

func (p *ProfileRepositoryImpl) GetProfile(header *Header) (map[string]any, error) {
	route := fmt.Sprintf("%s/account/profile", infra.AccountUrl)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Fatalf("failed create request | err : %v", err)
		return nil, err
	}

	request.Header.Set("User-ID", header.UserID)
	request.Header.Set("App-ID", header.AppID)
	request.Header.Set("Profile-ID", header.ProfileID)
	request.Header.Set("Authorization", header.Authorization)
	request.Header.Set("X-Api-Key", header.ApiKey)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("failed cliend do | err : %v", err)
		return nil, err
	}
	defer func() {
		if errBody := response.Body.Close(); errBody != nil {
			log.Fatalf("failed close response body | err : %v", err)
		}
	}()

	resp := FetchResponse(response)

	return resp, nil
}
