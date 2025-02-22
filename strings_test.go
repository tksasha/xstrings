package xstrings_test

import (
	"testing"

	strings "github.com/tksasha/xstrings"
	"gotest.tools/v3/assert"
)

func TestToSnakeCase(t *testing.T) {
	t.Run("when it is `HelloWorld`", func(t *testing.T) {
		subject := strings.ToSnakeCase("HelloWorld")

		assert.Equal(t, subject, "hello_world")
	})

	t.Run("when it is `CategoryID`", func(t *testing.T) {
		subject := strings.ToSnakeCase("CategoryID")

		assert.Equal(t, subject, "category_id")
	})

	t.Run("when it is `СлаваУкраїні`", func(t *testing.T) {
		subject := strings.ToSnakeCase("СлаваУкраїні")

		assert.Equal(t, subject, "слава_україні")
	})
}
