package service

import (
	"github.com/adlio/harvest"
	pass "github.com/manlycode/go-to-work/password"
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

func (hs *HarvestService) Name() string {
	return "Harvest"
}

func (hs HarvestService) FullName() string {
	return "GoToWork::Harvest"
}

func NewHarvestService() (harvestService *HarvestService) {
	return &HarvestService{}
}

// NewHarvestService creates a HarvestService instance
func (hs *HarvestService) SignIn(subdomain string, email string, password string) error {
	api := harvest.NewBasicAuthAPI(subdomain, email, password)
	res := WhoAmIResponse{}

	// Get the user
	err := api.Get(
		"/account/who_am_i",
		harvest.Defaults(),
		&res,
	)

	if err == nil {
		hs.User = res.User
		pw := pass.NewPassword(*hs)

		err = pw.Set(password)
	}

	return err
}
