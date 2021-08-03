package config

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/viper"
)

// Config stores all  the configuration of the application
type Config struct {
	AccessKey string `mapstructure:"ACCESS_KEY"`
	SecretKey string `mapstructure:"SECRET_KEY"`
	Region    string `mapstructure:"REGION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return
}

func BuildSession() *session.Session {
	// . means current folder
	config, err := LoadConfig("C:/Users/dell/go/src/github.com/amalikh/pms-primary-integration-service")

	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	sessionConfig := &aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
	}

	sess, err := session.NewSession(sessionConfig)
	if err != nil {
		fmt.Println("error", err)

	}
	return sess
}
