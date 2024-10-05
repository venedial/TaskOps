package config

type Configuration struct {
	Environment string
	Log         LogConfig
	Pgdb        PostgresDbConfig
}

type LogConfig struct {
	Stdout StdoutLogConfig
	File   *FileLogConfig
}

type StdoutLogConfig struct {
	Level string
}

type FileLogConfig struct {
	Level string
	Path  *string
}

type PostgresDbConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Sslmode  string
}
