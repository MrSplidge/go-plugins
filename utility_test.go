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

func TestMin(t *testing.T) {
	// Test case 1: Testing with integers
	resultInt := min(5, 3)
	if resultInt != 3 {
		t.Errorf("Expected min(5, 3) to be 3, but got %d", resultInt)
	}

	// Test case 1: Testing with integers
	resultInt2 := min(3, 3)
	if resultInt2 != 3 {
		t.Errorf("Expected min(3, 3) to be 3, but got %d", resultInt2)
	}

	// Test case 2: Testing with floats
	resultFloat := min(3.14, 2.71)
	if resultFloat != 2.71 {
		t.Errorf("Expected min(3.14, 2.71) to be 2.71, but got %f", resultFloat)
	}
}

func TestMax(t *testing.T) {
	resultInt := max(5, 2)
	if resultInt != 5 {
		t.Errorf("Expected max(5, 2) to be 5, but got %d", resultInt)
	}

	resultInt2 := max(3, 3)
	if resultInt2 != 3 {
		t.Errorf("Expected max(3, 3) to be 3, but got %d", resultInt2)
	}

	resultFloat := max(3.14, 2.71)
	if resultFloat != 3.14 {
		t.Errorf("Expected max(3.14, 2.71) to be 3.14, but got %f", resultFloat)
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
