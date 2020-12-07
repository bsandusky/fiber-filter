package filter

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestFilterRequest(t *testing.T) {

	app := fiber.New()

	app.Use(New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Passed filter")
	})

	t.Run("it should run filter function on all inbound requests", func(t *testing.T) {
		expected := "Passed filter"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp, _ := app.Test(req)
		body, _ := ioutil.ReadAll(resp.Body)

		utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode)
		utils.AssertEqual(t, expected, string(body))
	})
}

func TestBypassFilterRequest(t *testing.T) {

	app := fiber.New()

	app.Use(New(Config{
		Next: func(*fiber.Ctx) bool { return true },
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Passed filter")
	})

	t.Run("it should not run filter when Next is true in config", func(t *testing.T) {
		expected := "Passed filter"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp, _ := app.Test(req)
		body, _ := ioutil.ReadAll(resp.Body)

		utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode)
		utils.AssertEqual(t, expected, string(body))
	})
}

func TestFilterByIP(t *testing.T) {

	app := fiber.New()

	app.Use(New(Config{
		IPFilters: []string{"127.0.0.1", "0.0.0.0"},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Passed filter")
	})

	t.Run("it should filter requests by IP when match", func(t *testing.T) {
		expected := "request IP filtered"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp, _ := app.Test(req)
		body, _ := ioutil.ReadAll(resp.Body)

		utils.AssertEqual(t, fiber.StatusForbidden, resp.StatusCode)
		utils.AssertEqual(t, expected, string(body))
	})
}

func TestInvalidIPFilterString(t *testing.T) {

	app := fiber.New()

	app.Use(New(Config{
		IPFilters: []string{"*"},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Passed filter")
	})

	t.Run("it should throw error if string passed to filter is not compiled by regexp", func(t *testing.T) {
		expected := "cannot compile malformed filter string"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp, _ := app.Test(req)
		body, _ := ioutil.ReadAll(resp.Body)

		utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode)
		utils.AssertEqual(t, expected, string(body))
	})
}

func TestFilterByUserAgent(t *testing.T) {

	app := fiber.New()

	app.Use(New(Config{
		UserAgentFilters: []string{"PostmanRuntime*"},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Passed filter")
	})

	t.Run("it should filter requests by user agent when match", func(t *testing.T) {
		expected := "request user agent filtered"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("User-Agent", "PostmanRuntime/7.26.8")

		resp, _ := app.Test(req)
		body, _ := ioutil.ReadAll(resp.Body)

		utils.AssertEqual(t, fiber.StatusForbidden, resp.StatusCode)
		utils.AssertEqual(t, expected, string(body))
	})
}

func TestInvalidUserAgentFilterString(t *testing.T) {

	app := fiber.New()

	app.Use(New(Config{
		UserAgentFilters: []string{"*"},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Passed filter")
	})

	t.Run("it should throw error if string passed to filter is not compiled by regexp", func(t *testing.T) {
		expected := "cannot compile malformed filter string"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp, _ := app.Test(req)
		body, _ := ioutil.ReadAll(resp.Body)

		utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode)
		utils.AssertEqual(t, expected, string(body))
	})
}

func TestFilterXHRRequest(t *testing.T) {

	app := fiber.New()

	app.Use(New(Config{
		DisallowXHRRequest: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Passed filter")
	})

	t.Run("it should filter XHR requests when value set to true", func(t *testing.T) {
		expected := "XHR request filtered"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Add(fiber.HeaderXRequestedWith, "xmlhttprequest")

		resp, _ := app.Test(req)
		body, _ := ioutil.ReadAll(resp.Body)

		utils.AssertEqual(t, fiber.StatusForbidden, resp.StatusCode)
		utils.AssertEqual(t, expected, string(body))
	})
}

func TestFilterInsecureRequest(t *testing.T) {

	app := fiber.New()

	app.Use(New(Config{
		DisallowInsecure: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Passed filter")
	})

	t.Run("it should filter insecure requests when value set to true", func(t *testing.T) {
		expected := "insecure request filtered"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp, _ := app.Test(req)
		body, _ := ioutil.ReadAll(resp.Body)

		utils.AssertEqual(t, fiber.StatusForbidden, resp.StatusCode)
		utils.AssertEqual(t, expected, string(body))
	})
}
