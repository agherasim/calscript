package runtime

import "testing"

func TestListenerGetRule(t *testing.T) {
	rt, r := GetTestRule()
	rr, _ := NewRuleRegistry()
	rr.Add(rt, r)
	length, e := rr.Len()
	if length != 1 {
		t.Fatalf("unexpected length: expecting 1 got %d", length)
	}
	l, e := NewListener(rr)

	r, e = l.GetRule("test")
	if e != nil {
		t.Fatalf("failed to get rule from Listener Rule Registry")
	}

}
