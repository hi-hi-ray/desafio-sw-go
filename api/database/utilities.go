package database

import (
	"time"
)

func (configs *DatabaseConnection) CreateTimeout() time.Duration {
	return time.Duration(configs.Timeout) * time.Second
}
