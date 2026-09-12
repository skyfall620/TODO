package req

import "github.com/go-playground/validator/v10"

func IsValid[T any](payload T) error {
	validator := validator.New()

	if err := validator.Struct(payload); err != nil {
		return err
	}

	return nil
}
