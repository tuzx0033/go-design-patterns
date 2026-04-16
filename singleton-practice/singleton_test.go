package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter := GetInstance()

	if counter == nil {
		t.Error("expected pointer to Singleton after calling GetInstance, not nil")
	}

	expectedCounter := counter

	currentCount := counter.AddOne()
	if currentCount != 1 {
		t.Errorf("after first AddOne(), count must be 1 but got %d", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("expected same instance in counter, but got a different instance")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("after second AddOne(), count must be 2 but got %d", currentCount)
	}
}
