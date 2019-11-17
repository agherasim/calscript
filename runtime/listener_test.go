package runtime

import "testing"

func TestListenerGetRule(t *testing.T) {
	rt, r := GetTestRule()
	rr, _ := NewRuleRegistry()
	rr.Add(rt, r)
	length, err := rr.Len()
	if length != 1 {
		t.Fatalf("unexpected length: expecting 1 got %d", length)
	}
	l, err := NewCalscriptListener(rr)

	r, err = l.GetRule("test")
	if err != nil {
		t.Fatalf("failed to get rule from Listener Rule Registry %s", err)
	}

}

func TestParseShouldError(t *testing.T) {
	pp := getCalscriptParser()
	_, err := pp.Parse("test")

	if err.Error() != "line 1:0 missing 'calscript' at 'test'" {
		t.Fatalf("parser should fail with invalid script %s", err)
	}
}

func TestParseShouldSucceed(t *testing.T) {
	script := `
	calscript test_script
		on 2019/12/06 "I have an ëvent scheduled"
	`
	pp := getCalscriptParser()
	_, err := pp.Parse(script)

	if err != nil {
		t.Fatalf("parse should not fail with valid script %s", err)
	}
}

func getCalscriptParser() *CalscriptParser {
	rr, _ := NewRuleRegistry()
	rt, r := GetTestRule()
	rr.Add(rt, r)
	l, _ := NewCalscriptListener(rr)
	pp, _ := NewCalscriptParser(l)
	return pp
}
