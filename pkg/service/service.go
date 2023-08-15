package service

import (
	"fmt"
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CountDays() string {
	jan := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	t := jan.Sub(time.Now()).Hours() / 24
	return fmt.Sprintf("%d day(s) until 01 January 2025", int(t))
}
