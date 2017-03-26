package password

import (
	"github.com/tmc/keyring"
)

// Password contains information for writing a password to a keychain
type Password struct {
	ServiceName string
	UserName    string
}

type pwService interface {
	FullName() string
}

// NewPassword returns a password given a service
func NewPassword(s pwService) (pw *Password) {
	return &Password{
		ServiceName: s.FullName(),
	}
}

// Get returns the value for the given password
func (p *Password) Get() (string, error) {
	pw, err := keyring.Get(p.ServiceName, p.UserName)
	return pw, err
}

// Set returns the value for the given password
func (p *Password) Set(newPassword string) error {
	err := keyring.Set(p.ServiceName, p.UserName, newPassword)
	return err
}
