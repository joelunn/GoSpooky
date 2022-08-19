package config

// Configurations exported
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
	//EXAMPLE_PATH string
	//EXAMPLE_VAR  string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBHostname string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
}

func (r DatabaseConfigurations) GetConnectString() string {
	return r.DBUsername + ":" + r.DBPassword + "@tcp(" + r.DBHostname + ":" + r.DBPort + ")/" + r.DBName
}
