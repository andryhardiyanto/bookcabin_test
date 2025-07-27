package validators

import (
	"backend/pkg/datatype"
	"backend/pkg/fibers"
	"reflect"
	"strings"

	goValidator "github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type validator struct {
	validate *goValidator.Validate
}

type Validator interface {
	Validate(s any) *fibers.Error
}

func NewValidator() (Validator, error) {
	validation := goValidator.New()
	err := validation.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return nil, fibers.NewError(500, err.Error(), "internal_server_error")
	}
	var n = 2
	validation.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", n)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	err = validation.RegisterValidation("date", func(fl goValidator.FieldLevel) bool {
		field := fl.Field()
		if field.Kind() != reflect.Struct {
			return false
		}

		date, ok := field.Interface().(datatype.Date)
		if !ok {
			return false
		}

		return date.Valid
	})
	if err != nil {
		return nil, fibers.NewError(500, err.Error(), "internal_server_error")
	}

	return &validator{
		validate: validation,
	}, nil
}

func (v *validator) Validate(s any) *fibers.Error {
	err := v.validate.Struct(s)
	if err != nil {
		if errs, ok := err.(goValidator.ValidationErrors); ok {
			multipleErrors := make([]fibers.Violation, 0)
			for _, e := range errs {
				if strings.ToLower(e.Tag()) == "omitempty" {
					continue
				}

				field := e.Field()
				message := e.Error()

				multipleErrors = append(multipleErrors, fibers.Violation{
					Field:   field,
					Message: message,
				})
			}
			return fibers.NewViolationErrors(multipleErrors)
		}
	}

	return nil
}
