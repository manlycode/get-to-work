package service

import (
	"github.com/adlio/harvest"
	pass "github.com/manlycode/get-to-work/password"
)

// HarvestService defines a harvest service
type HarvestService struct {
	Service
	API  *harvest.API
	User *harvest.User
}

type WhoAmIResponse struct {
	User *harvest.User `json:"user"`
}

func NewHarvestService() (harvestService *HarvestService) {
	return &HarvestService{
		Service: Service{Name: "Harvest"},
	}
}

// NewHarvestService creates a HarvestService instance
func (hs *HarvestService) SignIn(subdomain string, email string, password string) (service HarvestService, err error) {
	api := harvest.NewBasicAuthAPI(subdomain, email, password)
	res := WhoAmIResponse{}
	srv := HarvestService{
		Service: Service{Name: "Harvest"},
		API:     api,
	}

	// Get the user
	err = api.Get(
		"/account/who_am_i",
		harvest.Defaults(),
		&res,
	)

	if err == nil {
		srv.User = res.User
		pw := pass.Password{
			ServiceName: srv.FullName(),
			UserName:    email,
		}

		err = pw.Set(password)
	}

	return srv, err
}
