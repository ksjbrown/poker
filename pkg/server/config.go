package server

// config object for server settings
type ServerConfig struct {
	Db DatabaseConfig
}

// DatabaseConfig defines values related to the backend database connection
type DatabaseConfig struct {
	Type DatabaseType
	Args map[DatabaseConfigOption]interface{}
}

type DatabaseType string

const (
	SQLITE DatabaseType = "sqlite"
)

type DatabaseConfigOption string
