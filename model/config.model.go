package model

// Config struct holding all sub details
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Log      LogConfig
}

// ServerConfig ...
type ServerConfig struct {
	Port int
}

// DatabaseConfig ...
type DatabaseConfig struct {
	Host     string
	Port     int
	DbName   string
	Username string
	Password string
}

// LogConfig ...
type LogConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}
