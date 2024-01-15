package main

import "fmt"

type Config struct {
	Address     string
	Port        uint16
	LogLevel    string
	Environment string
}

type ConfigOption func(Config) Config

func WithAddress(address string) ConfigOption {
	return func(config Config) Config {
		config.Address = address
		return config
	}
}

func WithPort(port uint16) ConfigOption {
	return func(config Config) Config {
		config.Port = port
		return config
	}
}

func WithLogLevel(logLevel string) ConfigOption {
	return func(config Config) Config {
		config.LogLevel = logLevel
		return config
	}
}

func WithEnvironment(environment string) ConfigOption {
	return func(config Config) Config {
		config.Environment = environment
		return config
	}
}

func CreateDefaultConfig() Config {
	return Config{LogLevel: "debug"}
}

func CreateConfig(options ...ConfigOption) Config {
	config := CreateDefaultConfig()

	for _, option := range options {
		config = option(config)
	}

	return config
}

func main() {
	config := CreateConfig(WithAddress("localhost"), WithPort(3000), WithEnvironment("development"))
	fmt.Println(config)
}
