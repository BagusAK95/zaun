package config

//Server : server configuration struct
type Server struct {
	Mode            string
	Addr            string
	Environment     string
	LogDuration     int
	ShutdownTimeout int
	BaseURL         string
	ClientURL       string
	Name            string
	Version         string
}
