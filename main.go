package main

import (
	"app/config"
	"app/database"
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	//database.Connect(AppConfig.ConnectionString)
	//database.Migrate()

	// Set undefined variables
	//viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Database server connection string is \t", configuration.Database.GetConnectString())

	database.Connect(configuration.Database.GetConnectString())
	database.InitialiseDB()

	// Reading variables using the model
	//fmt.Println("Reading variables using the model..")
	//fmt.Println("Database server is \t", configuration.Database.DBHostname)
	//fmt.Println("Database port is \t", configuration.Database.DBPort)
	//fmt.Println("Database name is \t", configuration.Database.DBName)
	//fmt.Println("Database user is \t", configuration.Database.DBUser)
	//fmt.Println("Database user is \t", configuration.Database.DBPassword)

	//fmt.Println("Front End Port is\t\t", configuration.Server.Port)
	//fmt.Println("---------------------\t")
	//fmt.Println("Database server connection string is \t", configuration.Database.GetConnectString())

	//fmt.Println("EXAMPLE_PATH is\t", configuration.EXAMPLE_PATH)
	//fmt.Println("EXAMPLE_VAR is\t", configuration.EXAMPLE_VAR)

	// Reading variables without using the model
	//fmt.Println("\nReading variables without using the model..")
	//fmt.Println("Database is\t", viper.GetString("database.dbname"))
	//fmt.Println("Port is\t\t", viper.GetInt("database.port"))
	//fmt.Println("EXAMPLE_PATH is\t", viper.GetString("EXAMPLE_PATH"))
	//fmt.Println("EXAMPLE_VAR is\t", viper.GetString("EXAMPLE_VAR"))

	/*
		e := echo.New()

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.GET("/", func(c echo.Context) error {
			return c.HTML(http.StatusOK, "Go Spooky!")
		})

		e.GET("/ping", func(c echo.Context) error {
			return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
		})

		httpPort := os.Getenv("HTTP_PORT")
		if httpPort == "" {
			httpPort = "8080"
		}

		e.Logger.Fatal(e.Start(":" + httpPort))
	*/
}
