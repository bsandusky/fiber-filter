package filter

import (
	"testing"
)

func TestFilterRequestByIP(t *testing.T) {

	t.Run("it should return true when no match is found -- default filters are nil", func(t *testing.T) {
		requestor := "127.0.0.1"

		allowed, _ := filter(requestor, nil)

		if !allowed {
			t.Error("filtered valid request")
		}
	})

	t.Run("it should return true when no match is found -- filters exist with no match", func(t *testing.T) {
		filters := []string{"127.0.0.2"}
		requestor := "127.0.0.1"

		allowed, _ := filter(requestor, filters)

		if !allowed {
			t.Error("filtered valid request")
		}
	})

	t.Run("it should return false when a match is found -- string given", func(t *testing.T) {
		filters := []string{"127.0.0.1"}
		requestor := "127.0.0.1"

		allowed, _ := filter(requestor, filters)

		if allowed {
			t.Error("did not filter invalid request")
		}
	})

	t.Run("it should return false when a match is found -- regex given", func(t *testing.T) {
		filters := []string{"127.0.0.*"}

		requestor := "127.0.0.1"
		allowed, _ := filter(requestor, filters)

		if allowed {
			t.Error("did not filter invalid request")
		}

		requestor = "127.0.0.25"
		allowed, _ = filter(requestor, filters)

		if allowed {
			t.Error("did not filter invalid request")
		}
	})

	t.Run("it should throw an error when string passed in not compiled by regexp lib", func(t *testing.T) {

		expected := "cannot compile malformed filter string"
		filters := []string{"*"}
		requestor := "127.0.0.1"
		_, err := filter(requestor, filters)

		if err == nil {
			t.Errorf("did not receive expected regexp error: %v", expected)
		}
	})
}
