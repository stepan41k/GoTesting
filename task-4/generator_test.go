package task4

import (
	"testing"
	"testing/quick"
)

func TestGenerateNumber(t *testing.T) {
	f := func(n uint8) bool {

		res := GenerateNumber(int(n+1))

		return res <= int(n) 
	} 

	if err := quick.Check(f, nil); err != nil {
		t.Errorf("expected failure: %s", err.Error())
	}
}
func TestGenerateString(t *testing.T) {
	f := func(n uint8) bool {
		length := int(n % 100)

		s := GenerateString(length)
		return len(s) == length
	}
	
	if err := quick.Check(f, nil); err != nil {
		t.Errorf("expected failure: %s", err.Error())
	}
}