package filter

import (
	"github.com/gofiber/fiber/v2"
)

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// IPFilters is used to define string and regex values for filters applied to
	// ctx.IP() of an inbound request.
	//
	// Optional. Default: nil
	IPFilters []string

	// UserAgentFilters is used to define string and regex values for filters
	// applied to string(c.Request().Header.UserAgent()) of an inbound request.
	//
	// Optional. Default: nil
	UserAgentFilters []string

	// DisallowXHRRequest is used to filter out requests coming from XHR client
	// libraries like jQuery
	//
	// Optional. Default: false
	DisallowXHRRequest bool

	// DisallowInsecure is used to filter out requests coming non-TLS connections
	//
	// Optional. Default: false
	DisallowInsecure bool
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:               nil,
	IPFilters:          nil,
	UserAgentFilters:   nil,
	DisallowXHRRequest: false,
	DisallowInsecure:   false,
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values for missing fields in overrides
	if cfg.Next == nil {
		cfg.Next = ConfigDefault.Next
	}

	if cfg.IPFilters == nil {
		cfg.IPFilters = ConfigDefault.IPFilters
	}

	if cfg.UserAgentFilters == nil {
		cfg.UserAgentFilters = ConfigDefault.UserAgentFilters
	}

	return cfg
}
