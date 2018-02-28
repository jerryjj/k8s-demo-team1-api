package main

import (
	"os"

	logging "github.com/op/go-logging"
)

// Initializes our local logger
func SetupLocalLogger(logModule string) *logging.Logger {
	var format = logging.MustStringFormatter("%{color}%{time:15:04:05.000} " +
		"%{shortfunc} ▶ %{level} " +
		"%{color:reset} %{message}")
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(formatter)

	log := logging.MustGetLogger(logModule)
	logging.SetLevel(logging.DEBUG, logModule)

	// Compensate for all the wrapping layers around the logger
	log.ExtraCalldepth = 3

	log.Debug("Debug logging enabled.")

	return log
}

// Reads an environment variable value or panics if not found
func MustGetenv(varname string, default_value string) string {
	log.Debugf("Reading environment variable %v", varname)

	value := os.Getenv(varname)
	if value == "" {
		if default_value == "" {
			log.Fatalf("Missing env variable %v", varname)
		}
		value = default_value
	}

	return value
}

var (
	log = SetupLocalLogger("payments")
)

const (
	// Environment variable for payments address
	paymentsAddressEnvVar = "PAYMENTS_ADDRESS"
)

func main() {
	// Read things out of the environment
	paymentsAddress := MustGetenv(paymentsAddressEnvVar, "localhost:50051")
	mustCreatePaymentsClient(paymentsAddress)

	mustRunRESTServer(8080)
}
