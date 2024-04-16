package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "hello"}
	hello2 := &String{Value: "hello"}
	diff1 := &String{Value: "diff"}
	diff2 := &String{Value: "diff"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("hello1 and hello2 should have the same hash key")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("diff1 and diff2 should have the same hash key")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("hello1 and diff1 should have different hash key")
	}
}
