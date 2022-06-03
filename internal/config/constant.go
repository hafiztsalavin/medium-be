package config

import "time"

const (
	// TimeoutSecond used for timeout on handler decorators WithTimeout
	TimeoutSecond = 5
	// ReadTimeout is the maximum duration for reading the entire request, including body
	ReadTimeout = 310 * time.Second
	// WriteTimeout is the maximun duration before timing out writes of response
	WriteTimeout = 310 * time.Second
)
