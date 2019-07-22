package config

import "time"

//Redis : redis configuration struct
type Redis struct {
	Addr     string
	Password string
	Database int
	TTL      time.Duration
}
