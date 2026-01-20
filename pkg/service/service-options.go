package service

import "github.com/TamasGorgics/gomag/pkg/logx"

type ServiceOptions func(*Service)

func WithLogger(logger logx.Logger) ServiceOptions {
	return func(s *Service) {
		s.logger = logger
	}
}
