package models

import (
	"fmt"

	"github.com/loopcontext/auth-api-go/src/utils/consts"
	"github.com/loopcontext/auth-api-go/src/utils"
)

type RegisterInput struct {
	LoginInput
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (li *RegisterInput) Validate() error {
	if li.Email == "" {
		return fmt.Errorf(consts.ErrStrNotEmpty, "email")
	}
	if li.FirstName == "" {
		return fmt.Errorf(consts.ErrStrNotEmpty, "firstName")
	}
	if li.LastName == "" {
		return fmt.Errorf(consts.ErrStrNotEmpty, "lastName")
	}
	if li.Password == "" {
		return fmt.Errorf(consts.ErrStrNotEmpty, "password")
	}
	if err := utils.EmailCheck(li.Email); err != nil {
		return err
	}

	return nil
}
