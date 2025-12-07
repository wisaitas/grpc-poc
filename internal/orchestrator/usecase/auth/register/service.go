package register

import "context"

type Service interface {
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Register(ctx context.Context, req RegisterRequest) error {
	return nil
}
