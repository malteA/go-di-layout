package application

import (
	"fmt"
	"math/rand"
	"strconv"
)

type TransactionLog interface {
	Log(id, message string)
}

type CreditCardProcessor interface {
	Auth(amount float64) error
	Capture(amount float64) error
}

type Service struct {
	logger    TransactionLog
	processor CreditCardProcessor
}

func (s *Service) Inject(logger TransactionLog, processor CreditCardProcessor) *Service {
	s.logger = logger
	s.processor = processor
	return s
}

func (s *Service) MakeTransaction(amount float64, message string) error {
	id := strconv.Itoa(rand.Int())

	s.logger.Log(id, fmt.Sprintf("Starting transaction %q", message))

	s.logger.Log(id, "Try to Auth")
	if err := s.processor.Auth(amount); err != nil {
		s.logger.Log(id, "Auth failed")
		return err
	}

	s.logger.Log(id, "Try to Capture")
	if err := s.processor.Capture(amount); err != nil {
		s.logger.Log(id, "Capture failed")
		return err
	}

	s.logger.Log(id, "Transaction successful")
	return nil
}
