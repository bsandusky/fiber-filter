# fiber-filter

The package adds support for request filtering to the excellent [Fiber](https://gofiber.io/) framework. Once applied, `fiber-filter` middleware will use regexp to check for matches for values passed into the config object in order to filter inbound requests to a route, a group, or an entire app. Filters can be applied to inbound IP addresses, User Agent strings, XHR request status, and TLS request status.

## Table of Contents

- [Signatures](#signatures)
- [Examples](#examples)
- [Config](#config)
- [Default Config](#default-config)

## Signatures

```go
func New(config ...Config) fiber.Handler
```

## Examples

### Basic Usage as Middleware

```go
    app := fiber.New()

    // Initialize default config
    app.Use(filter.New())

    // Or extend default config with customizations
    app.Use(filter.New(filter.Config{
        IPFilters: []string{"127.0.0.*"},
        AllowXHRRequest: false,
    }))

    // Filtered routes
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Passed filtering rules!")
    })

```

## Config

```go
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
```

## Default Config

```go
// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:               nil,
	IPFilters:          nil,
	UserAgentFilters:   nil,
	DisallowXHRRequest: false,
	DisallowInsecure:   false,
}
```
