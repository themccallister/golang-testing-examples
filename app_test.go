package main

import "testing"

// TESTS

func TestCanSubtractOneFromANumber(test *testing.T) {
	if integer := subtractByOne(2); integer != 1 {
		test.Error("The returned value was not what we expected")
	} else {
		test.Log("Can subtract a number from any int!")
	}
}

func TestCanAddStringsToAnArray(test *testing.T) {
	var arr = addStringsToArray("one", "two")

	if arr[0] != "one" {
		test.Error("Expected one, got ", arr[0], " instead...")
	} else {
		test.Log("Saw one!")
	}
	if arr[1] != "two" {
		test.Error("Expected two, got ", arr[1], " instead...")
	} else {
		test.Log("Saw two!")
	}
}

// BENCHMARKS

func BenchmarkSubtractingOneFromANumber(test *testing.B) {
	test.StopTimer()
	test.StartTimer()
	for i := 0; i < test.N; i++ {
		subtractByOne(2)
	}
}

func BenchmarkAddingStringToAnArray(test *testing.B) {
	test.StopTimer()
	test.StartTimer()
	for i := 0; i < test.N; i++ {
		addStringsToArray("one", "two")
	}
}
