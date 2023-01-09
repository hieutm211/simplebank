package api

import (
	"simplebank/util"

	"github.com/go-playground/validator/v10"
)

var currencyValidator validator.Func = func(fl validator.FieldLevel) bool {
	currency, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	return util.IsCurrencySupported(currency)
}
