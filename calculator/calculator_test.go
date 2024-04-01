package calculator_test

import (
	"calculatorProject/calculator"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDivide(t *testing.T) {
	expected := 2.0
	got := calculator.Divide(10.0, 5.0)

	if expected != got {
		t.Errorf("expected %.1f, got %.1f", expected, got)
	}

}

func TestDivideNegativeDivisor(t *testing.T) {
	expected := -2.0
	got := calculator.Divide(10.0, -5.0)

	// if got != expected {
	// 	t.Fatalf("expected %.1f, got %.1f", expected, got)
	// }

	assert.Equal(t, expected, got)
}
