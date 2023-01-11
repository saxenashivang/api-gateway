package constants

import "time"

const (
	RateLimitPeriod          = 15 * time.Minute
	RateLimitRequestsPerUser = int64(10)
)
