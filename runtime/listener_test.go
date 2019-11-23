package runtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListenerGetRule(t *testing.T) {
	rt, r := getTestRule("test")
	rr, _ := NewRuleRegistry()
	rr.Add(rt, r)
	length, err := rr.Len()
	assert.Equal(t, 1, length)
	assert.Nil(t, err)

	l, err := NewListener(rr)
	assert.Nil(t, err)
	r1, err := l.GetRule("test")
	assert.Nil(t, err)
	assert.ObjectsAreEqual(r, r1)
}
