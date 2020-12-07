package filter

import (
	"github.com/gofiber/fiber/v2"
)

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {

	// Set default config
	cfg := configDefault(config...)

	// Return new handler
	return func(c *fiber.Ctx) error {

		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Run filters on request
		if err := filterRequest(cfg, c); err != nil {
			return err
		}

		// Continue stack
		return c.Next()
	}
}

func filterRequest(cfg Config, c *fiber.Ctx) error {

	// Checks based on boolean values
	if cfg.DisallowXHRRequest == true && c.XHR() == true {
		return fiber.NewError(fiber.StatusForbidden, "XHR request filtered")
	}

	if cfg.DisallowInsecure == true && c.Secure() == false {
		return fiber.NewError(fiber.StatusForbidden, "insecure request filtered")
	}

	// Checks based on pattern matching
	allowed, err := filter(c.IP(), cfg.IPFilters)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if !allowed {
		return fiber.NewError(fiber.StatusForbidden, "request IP filtered")
	}

	allowed, err = filter(string(c.Request().Header.UserAgent()), cfg.UserAgentFilters)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if !allowed {
		return fiber.NewError(fiber.StatusForbidden, "request user agent filtered")
	}

	return nil
}
