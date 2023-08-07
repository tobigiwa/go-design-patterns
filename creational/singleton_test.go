package creational_test

import (
	"dsa/creational"
	"fmt"
	"testing"
)

func TestGetSingletonInstance(t *testing.T) {
	fmt.Println("Test started...")

	counter1 := creational.GetSingletonInstance()
	if counter1 == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}
	expectedCounter := counter1
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}

	counter2 := creational.GetSingletonInstance()
	if counter2 != expectedCounter {
		t.Error("Expected same instance in counter2 but it got a different instance")
	}
	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the currentncount must be 2 but was %d\n", currentCount)
	}
}
