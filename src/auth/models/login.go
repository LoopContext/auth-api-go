package models

import (
	"fmt"

	"github.com/loopcontext/auth-api-go/src/utils/consts"
	"github.com/loopcontext/auth-api-go/src/utils"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (li *LoginInput) Validate() error {
	if li.Email == "" {
		return fmt.Errorf(consts.ErrStrNotEmpty, "email")
	}
	if li.Password == "" {
		return fmt.Errorf(consts.ErrStrNotEmpty, "password")
	}
	if err := utils.EmailCheck(li.Email); err != nil {
		return err
	}

	return nil
}
