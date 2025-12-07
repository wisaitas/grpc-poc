package initial

import "github.com/wisaitas/grpc-poc/pkg/validatorx"

type SDK struct {
	Validatorx validatorx.Validator
}

func newSDK() *SDK {
	return &SDK{
		Validatorx: validatorx.NewValidator(),
	}
}
