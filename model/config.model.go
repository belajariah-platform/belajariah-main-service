package model

// Config struct holding all sub details
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Log      LogConfig
	Mail     MailConfig
	System   ServiceSystem
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
	SSLMode  string
	SSLRoot  string
}

// LogConfig ...
type LogConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

// MailConfig ...
type MailConfig struct {
	AuthEmail    string
	AuthPassword string
	SMTPHost     string
	SMTPPort     int
	SenderName   string
	CopyRight    string
}

type ServiceSystem struct {
	EmailSystem string
}
