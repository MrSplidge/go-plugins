package main

import (
	"reflect"
	"testing"
)

func TestPadStringToWidth(t *testing.T) {
	result := padStringToWidth("test", 10, '.')
	expected := "test......"
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestDedup(t *testing.T) {
	// Test case with duplicate elements
	input1 := []int{1, 2, 2, 3, 6, 6, 4, 4, 4, 5}
	expected1 := []int{1, 2, 3, 6, 4, 5}
	result1 := dedup(input1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Dedup test case 1 failed, got: %v, want: %v", result1, expected1)
	}

	// Test case with no duplicates
	input2 := []string{"apple", "banana", "cherry", "cherry"}
	expected2 := []string{"apple", "banana", "cherry"}
	result2 := dedup(input2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Dedup test case 2 failed, got: %v, want: %v", result2, expected2)
	}

	// Test case with empty input
	input3 := []float64{}
	expected3 := []float64{}
	result3 := dedup(input3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Dedup test case 3 failed, got: %v, want: %v", result3, expected3)
	}

	// Test case with all items the same
	input4 := []string{"a", "a", "a", "a"}
	expected4 := []string{"a"}
	result4 := dedup(input4)
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Dedup test case 4 failed, got: %v, want: %v", result4, expected4)
	}
}

func TestIterateOverCISortedMap(t *testing.T) {
	testMap := map[string]int{"z": 1, "a": 2, "B": 3}
	expectedOrder := []string{"a", "B", "z"}
	expectedValues := []int{2, 3, 1}

	var keys []string
	var values []int
	iterateOverCISortedMap(testMap, func(key string, value int) {
		keys = append(keys, key)
		values = append(values, value)
	})

	if !reflect.DeepEqual(keys, expectedOrder) {
		t.Errorf("Keys not sorted case-insensitively. Expected: %v, Got: %v", expectedOrder, keys)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Values not matching the sorted keys. Expected: %v, Got: %v", expectedValues, values)
	}
}
