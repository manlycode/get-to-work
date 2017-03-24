package harvest

import (
	"github.com/adlio/harvest"
)

// Service represents the api client and the logged in user for Harvest
type Service struct {
	API  *harvest.API
	User *harvest.User
}

type WhoAmIResponse struct {
	User *harvest.User `json:"user"`
}

func NewService(subdomain string, email string, password string) (service Service, err error) {
	api := harvest.NewBasicAuthAPI(subdomain, email, password)
	res := WhoAmIResponse{}
	err = api.Get("/account/who_am_i", harvest.Defaults(), &res)

	return Service{API: api, User: res.User}, err
}
