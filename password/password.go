package password

import (
	"github.com/tmc/keyring"
)

// Password contains information for writing a password to a keychain
type Password struct {
	ServiceName string
	UserName    string
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
