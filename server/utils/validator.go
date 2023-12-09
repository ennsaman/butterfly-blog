package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var Validator = new(ValidateUtil)

type ValidateUtil struct{}

var (
	validate *validator.Validate
)

// Validate 合法性检查
func (v *ValidateUtil) Validate(data any) string {
	ans := ""
	validate = validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(data)
	if err != nil {
		fmt.Println("validate err: ", err)
		for _, e := range err.(validator.ValidationErrors) {
			ans += fmt.Sprintf("Field %s, Error Tag %s, Error Param %s\n", e.Field(), e.Tag(), e.Param())
		}
	} else {
		fmt.Println("validate success")
	}
	return ans
}
