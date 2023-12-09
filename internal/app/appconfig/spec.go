package appconfig

import (
	"github.com/shaband/POS/internal/app/appcontext"
)

// ConfigSpec is the configuration specification.
type ConfigSpec struct {
	// ServiceListenAddress is the address that the Fiber HTTP server will listen on.
	ServiceListenAddress string `default:":3000" required:"true" split_words:"true"`

	// LogJSONStdout is the flag to enable JSON logging to stdout and disable logging to file.
	LogJSONStdout bool `default:"false" required:"true" split_words:"true"`

	// LogLevel is the log level. Valid values are: trace, debug, info, warn, error, fatal, panic.
	LogLevel ConfigLogLevel `default:"info" required:"true" split_words:"true"`
}

type Config struct {
	// ConfigSpec is the configuration specification injected to the config.
	ConfigSpec

	// AppContext is the application context
	AppContext appcontext.Ctx
}
