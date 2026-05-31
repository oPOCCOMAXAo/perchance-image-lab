package middleware

import "github.com/opoccomaxao/gopkg/pkg/services/tasks"

type Service struct {
	tasks *tasks.Service
}

func NewMiddleware(
	tasks *tasks.Service,
) *Service {
	return &Service{
		tasks: tasks,
	}
}
