package main

import "testing"

func TestIsDecreasing(t *testing.T) {
	if isDecreasing("111123") == true {
		t.Errorf("Number not decreasing")
	}

	if isDecreasing("223450") == false {
		t.Errorf("Number is decreasing pair of digits 50")
	}
}

func TestHasDouble(t *testing.T) {
	if withAtLeastOneDouble(getDoubledDigits("1222345")) != true {
		t.Errorf("Number has double, pair of 22")
	}

	if withAtLeastOneDouble(getDoubledDigits("123789")) == true {
		t.Errorf("Number does not have double")
	}
}
