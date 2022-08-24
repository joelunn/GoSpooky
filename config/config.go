package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *Configurations

// Configurations exported
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
	//EXAMPLE_PATH string
	//EXAMPLE_VAR  string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port string
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBHostname string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
}

func GetConnectString() string {
	return Config.Database.DBUsername + ":" + Config.Database.DBPassword + "@tcp(" + Config.Database.DBHostname + ":" + Config.Database.DBPort + ")/" + Config.Database.DBName
}

func GetServerPort() string {
	return Config.Server.Port
}

func LoadAppConfig() {

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
}
