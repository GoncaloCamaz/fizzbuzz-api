package app

// FizzBuzzServiceConfiguration holds configuration for the FizzBuzz service
type FizzBuzzServiceConfiguration struct {
	StatisticsRPCPath string
	ServiceHTTPPath   string
	ServiceName       string
}

// NewFizzBuzzServiceConfiguration creates a new FizzBuzzServiceConfiguration with custom values
func NewFizzBuzzServiceConfiguration(statisticsRPCPath, serviceHTTPPath,
	serviceName string) *FizzBuzzServiceConfiguration {
	return &FizzBuzzServiceConfiguration{
		StatisticsRPCPath: statisticsRPCPath,
		ServiceHTTPPath:   serviceHTTPPath,
		ServiceName:       serviceName,
	}
}
