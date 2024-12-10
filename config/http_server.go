package config

import (
	"fmt"
)

const (
	domain = "localhost"
	port   = 8080
)

func GetAddress() string {
	return fmt.Sprintf("%s:%d", domain, port)
}
