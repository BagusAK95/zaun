package config

//Logger : logger configuration struct
type Logger struct {
	Host     string
	Port     string
	Level    string
	Username string
	Password string
	Tags     []string
}
