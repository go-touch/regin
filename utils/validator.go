package utils

type ValidatorHandler struct {
}

var Validator *ValidatorHandler

func init() {
	Validator = &ValidatorHandler{}
}
