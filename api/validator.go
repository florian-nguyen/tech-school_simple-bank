package api

import (
	"github.com/florian-nguyen/tech-school_simple-bank/simple-bank/db/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// Check if the currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
