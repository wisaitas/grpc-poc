package initial

import "github.com/wisaitas/grpc-poc/pkg/validatorx"

type sdk struct {
	validatorx validatorx.Validator
}

func newSDK() *sdk {
	return &sdk{
		validatorx: validatorx.NewValidator(),
	}
}
