package main

import "fmt"

type Config struct {
	Address     string
	Port        uint16
	LogLevel    string
	Environment string
}

func (config *Config) WithAddress(address string) *Config {
	config.Address = address
	return config
}

func (config *Config) WithPort(port uint16) *Config {
	config.Port = port
	return config
}

func (config *Config) WithLogLevel(logLevel string) *Config {
	config.LogLevel = logLevel
	return config
}

func (config *Config) WithEnvironment(environment string) *Config {
	config.Environment = environment
	return config
}

func (config *Config) Build() Config {
	return *config
}

func DefaultConfig() *Config {
	return &Config{LogLevel: "debug"}
}

func main() {
	config := DefaultConfig().WithAddress("localhost").WithPort(3000).WithEnvironment("development").Build()
	fmt.Println(config)
}
