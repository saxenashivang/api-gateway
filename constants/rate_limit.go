package constants

import "time"

const (
	RateLimitPeriod          = 2 * time.Second
	RateLimitRequestsPerUser = int64(10)
)
