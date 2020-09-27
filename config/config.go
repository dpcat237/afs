package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

const (
	// ParameterNamePortHTTP contains parameter port http name
	ParameterNamePortHTTP = "portHttp"
	// PortHTTP is server port
	PortHTTP = 8080
)

var (
	portHttp = flag.Int(ParameterNamePortHTTP, PortHTTP, "HTTP server port")
)

// Config defines configuration parameters.
type Config struct {
	PortHTTP string
}

// LoadConfigData loads environment parameters.
func LoadData() (Config, error) {
	var cfg Config
	envPortHTTP := os.Getenv("AFS_PORT_HTTP")
	if envPortHTTP != "" {
		var value int

		value, err := strconv.Atoi(envPortHTTP)
		if err != nil {
			return Config{}, err
		}
		*portHttp = value
	}
	cfg.PortHTTP = fmt.Sprintf("%d", *portHttp)

	return cfg, nil
}
