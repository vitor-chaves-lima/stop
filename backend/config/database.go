package config

import "fmt"

type MongoDBConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func (c *MongoDBConfig) GenerateDSN() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin", c.Username, c.Password, c.Host, c.Port, c.Database)
}
