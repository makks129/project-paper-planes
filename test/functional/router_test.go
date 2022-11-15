package functional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStart(t *testing.T) {

	t.Run("returns reply, if reply exists", func(t *testing.T) {
		// TODO
		assert.Equal(t, 1, 1)
	})

	t.Run("returns message, if assigned message exists", func(t *testing.T) {
		// TODO
		assert.Equal(t, 1, 1)
	})

	t.Run("doesn't return message, if assigned message exists but it's already read", func(t *testing.T) {
		// TODO
		assert.Equal(t, 1, 1)
	})

	t.Run("returns message, if assigned-unread message doesn't exist and unassigned one exists", func(t *testing.T) {
		// TODO
		assert.Equal(t, 1, 1)
	})

	t.Run("doesn't return message, if neither assigned-unread or unassigned messages exist", func(t *testing.T) {
		// TODO
		assert.Equal(t, 1, 1)
	})

}
