package app

// StatisticsServiceConfiguration holds configuration for the Statistics service
type StatisticsServiceConfiguration struct {
	ServiceRPCPath  string
	ServiceHTTPPath string
	ServiceName     string
}

// NewStatisticsServiceConfiguration creates a new StatisticsServiceConfiguration with custom values
func NewStatisticsServiceConfiguration(rpcServer, httpServer, serviceName string) *StatisticsServiceConfiguration {
	return &StatisticsServiceConfiguration{
		ServiceRPCPath:  rpcServer,
		ServiceHTTPPath: httpServer,
		ServiceName:     serviceName,
	}
}
