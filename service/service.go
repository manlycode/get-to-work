package service

import (
	"fmt"
	"github.com/manlycode/get-to-work/password"
)

// Service defines a service used in get-to-work
type Service struct {
	Name string
}

func (s *Service) FullName() string {
	return fmt.Sprintf("GoToWork::%s", s.Name)
}

// Password builds a password object from the service
func (s *Service) Password() (pw *password.Password) {
	return &password.Password{
		ServiceName: s.FullName(),
	}
}
