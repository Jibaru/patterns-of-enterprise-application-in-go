package gateways

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jibaru/gateway/internal/models"
)

type UserGateway struct {
	baseURL string
}

func NewUserGateway(baseURL string) *UserGateway {
	return &UserGateway{baseURL: baseURL}
}

func (ug *UserGateway) GetUserByID(id int) (*models.User, error) {
	url := fmt.Sprintf("%s/users/%d", ug.baseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &user, nil
}
