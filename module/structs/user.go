package structs

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

type User struct {
	Id          int    `json:"id,omitempty" db:"id"`
	FirstName   string `json:"fName"  validate:"required" db:"first_name"`
	LastName    string `json:"lName"  validate:"required" db:"last_name"`
	Phone       string `json:"phone"  validate:"required" db:"phone"`
	GovId       string `json:"govId"  validate:"required" db:"gov_id"`
	Password    string `json:"password" db:"password"`
	Role        string `json:"role"  validate:"required" db:"role"`
	AccessToken string `json:"accessToken"  validate:"required" db:"access_token"`
}

func (c *User) Validate() error {

	validate := validator.New()

	// Validate the struct
	err := validate.Struct(c)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.New("validation failed")
		}

		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Field(), err.Tag())
		}
		return errors.New("validation errors: " + strings.Join(validationErrors, ", "))
	}
	return nil
}
