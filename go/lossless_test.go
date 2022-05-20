package main

import (
	c "tdd/stocks"
	"testing"
)

func TestCompressionLossLessFirst(t *testing.T) {
	actual := c.ExecuteCompressionLossLess("AAABBCDDDDDEFFGH")
	expected := "A3B2C1D5E1F2G1H1"

	if expected != actual {
		t.Errorf("Expected [%+v] Got: [%+v]", expected, actual)
	}
}

func TestCompressionLossLessSecond(t *testing.T) {
	actual := c.ExecuteCompressionLossLess("AABCCCCDDEFFFGGHHHH")
	expected := "A2B1C4D2E1F3G2H4"

	if expected != actual {
		t.Errorf("Expected [%+v] Got: [%+v]", expected, actual)
	}
}