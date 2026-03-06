package command

import (
	"github.com/brunobotter/site-sentinel/application"
	"github.com/brunobotter/site-sentinel/application/validator"
)

const maxPaginationLimit = 100

type CreateTenant struct {
	Name string
}

func (c *CreateTenant) Validate() error {
	v := validator.NewFieldValidatorControl()
	v.AddFieldValidator("name", c.Name, validator.Required())
	return application.NewValidationApplicationError(application.ValidationDomain, v.Error())
}

type ListTenant struct {
	Page  int
	Limit int
}

func (c *ListTenant) Validate() error {
	v := validator.NewFieldValidatorControl()
	v.AddFieldValidator("page", c.Page, validator.MinNumber(1))
	v.AddFieldValidator("limit", c.Limit, validator.MinNumber(1), validator.MaxNumber(maxPaginationLimit))
	return application.NewValidationApplicationError(application.ValidationDomain, v.Error())
}
