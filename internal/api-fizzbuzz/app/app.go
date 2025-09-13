/*
Package app holds utilities to start and manage the application
*/
package app

import (
	"fmt"
)

// FizzBuzzService represents the fizzbuzz service
type FizzBuzzService struct {
	configuration *FizzBuzzServiceConfiguration
}

// NewFizzBuzzService returns a new FizzBuzzService
func NewFizzBuzzService(conf *FizzBuzzServiceConfiguration) *FizzBuzzService {
	return &FizzBuzzService{
		configuration: conf,
	}
}

func (f *FizzBuzzService) StartFizzBuzzService() error {
	fmt.Println("Starting FizzBuzzService...", f.configuration)

	return nil
}
