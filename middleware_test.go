package Logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	t.Run("Add middleware", func(t *testing.T) {
		middlewares = nil

		AddMiddleware(&TestLogger{})

		assert.NotNil(t, middlewares)
		assert.Len(t, middlewares, 1)
	})

	t.Run("Middleware order", func(t *testing.T) {
		t.Skip("Not implemented yet")
	})
}
