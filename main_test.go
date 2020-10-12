package main

import "testing"

func TestCalculatePrice(t *testing.T) {
	result := CalculatePrice(2, 3)
	if result !=  10799.98 {
		t.Errorf("Test failed for Macbok Pro, Expected  10799.98, got %v", result)
	}

	result = CalculatePrice(3, 4)
	if result !=  394.20 {
		t.Errorf("Test failed for Alexa Speaker, Expected  394.20, got %v", result)
	}
}
