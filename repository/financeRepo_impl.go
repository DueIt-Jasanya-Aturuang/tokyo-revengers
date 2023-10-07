package repository

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/infra"
	"github.com/DueIt-Jasanya-Aturuang/tokyo-revengers/util"
)

type FinanceRepositoryImpl struct {
}

func (p *FinanceRepositoryImpl) GetPayment(header *Header, order, cursor string) (map[string]any, error) {
	route := fmt.Sprintf("%s/finance/payment?order=%s&cursor=%s", infra.FinanceUrl, order, cursor)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("User-ID", header.UserID)
	request.Header.Set("App-ID", header.AppID)
	request.Header.Set("Profile-ID", header.ProfileID)
	request.Header.Set("Authorization", header.Authorization)
	request.Header.Set("X-Api-Key", infra.FinanceKey)

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

func (p *FinanceRepositoryImpl) GetSpendingType(header *Header, order, cursor string, configValue string) (map[string]any, error) {
	var route string
	if configValue == "" {
		route = fmt.Sprintf("%s/finance/spending-type?order=%s&cursor=%s", infra.FinanceUrl, order, cursor)
	} else {
		route = fmt.Sprintf("%s/finance/spending-type/%s?order=%s&cursor=%s", infra.FinanceUrl, configValue, order, cursor)
	}

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("User-ID", header.UserID)
	request.Header.Set("App-ID", header.AppID)
	request.Header.Set("Profile-ID", header.ProfileID)
	request.Header.Set("Authorization", header.Authorization)
	request.Header.Set("X-Api-Key", infra.FinanceKey)

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

func (p *FinanceRepositoryImpl) GetDetailSpendingType(header *Header, id string) (map[string]any, error) {
	route := fmt.Sprintf("%s/finance/spending-type/detail/%s", infra.FinanceUrl, id)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("User-ID", header.UserID)
	request.Header.Set("App-ID", header.AppID)
	request.Header.Set("Profile-ID", header.ProfileID)
	request.Header.Set("Authorization", header.Authorization)
	request.Header.Set("X-Api-Key", infra.FinanceKey)

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
