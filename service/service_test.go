package service

import (
	"testing"
)

func TestAccountName(t *testing.T) {
	s := Service{Name: "Example"}
	acctName := s.AccountName()

	if acctName != "GoToWork::Example" {
		t.Errorf(
			"Expected acctName to be GotoWork::Example but was %s",
			acctName,
		)
	}
}
