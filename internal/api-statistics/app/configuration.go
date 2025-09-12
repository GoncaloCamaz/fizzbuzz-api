package app

import "fizzbuzz-api/pkg/database"

// StatisticsServiceConfiguration holds configuration for the Statistics service
type StatisticsServiceConfiguration struct {
	ServiceRPCPath  string
	ServiceHTTPPath string
	ServiceName     string
	DB              *database.Config
}

// NewStatisticsServiceConfiguration creates a new StatisticsServiceConfiguration with custom values
func NewStatisticsServiceConfiguration(rpcServer, httpServer, serviceName string,
	db *database.Config) *StatisticsServiceConfiguration {
	return &StatisticsServiceConfiguration{
		ServiceRPCPath:  rpcServer,
		ServiceHTTPPath: httpServer,
		ServiceName:     serviceName,
		DB:              database.DefaultDBConfig(),
	}
}
