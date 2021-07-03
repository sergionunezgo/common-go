package service

// Config defines the values that can be loaded from env vars or other config files.
type Config struct {
	Port     int
	LogLevel string
}

// Service defines the methods that are required to operate a web service.
type Service interface {
	Start() error
	Close() error
}
