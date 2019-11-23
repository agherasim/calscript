package runtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test error listener internal stack doesn't exceed capacity
func TestErrorListenerErrorStack(t *testing.T) {
	el := NewErrorListener(5)
	assert.Equal(t, cap(el.GetErrors()), 5)

	for i := 0; i < 6; i++ {
		el.AppendError(i, i, "error")
	}
	assert.Len(t, el.GetErrors(), 5)
}
