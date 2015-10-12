package welchttest

import (
	"math"
	"testing"
)

func TestStudentsCDF(t *testing.T) {
	// We should get 0 if dgf is 0
	expected := float64(0)
	actual := StudentsT_CDF(0, 0, true)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	// We should get 0.5 if t-score is 0
	expected = float64(0.5)
	actual = StudentsT_CDF(0, 10, true)
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	// The following tests try various combinations of the parameters whose correct
	// probability has been determined through external testing.
	expected = float64(0.1955011)
	actual = StudentsT_CDF(1, 3, true)
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(0.07797913)
	actual = StudentsT_CDF(4, 1, true)
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(0.92202087)
	actual = StudentsT_CDF(4, 1, false)
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(0.02410609)
	actual = StudentsT_CDF(2, 100, true)
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(0.9758939)
	actual = StudentsT_CDF(2, 100, false)
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(0.9758939)
	actual = StudentsT_CDF(-2, 100, true)
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(0.02410609)
	actual = StudentsT_CDF(-2, 100, false)
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
func TestGetConfidence(t *testing.T) {
	// We should get 0 confidence if either count is less than 1.
	expected := float64(0)
	actual := GetConfidence(0, 12, float64(0.5), float64(0.25), float64(0.2), float64(0.2))
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(0)
	actual = GetConfidence(12, 0, float64(0.5), float64(0.25), float64(0.2), float64(0.2))
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	// We should also get 0 confidence if the mean of x is less than or equal to the mean of
	// y.
	expected = float64(0)
	actual = GetConfidence(12, 12, float64(0.25), float64(0.25), float64(0.2), float64(0.2))
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(0)
	actual = GetConfidence(12, 12, float64(0.24), float64(0.25), float64(0.2), float64(0.2))
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	// The following tests try various combinations of the parameters whose resulting expected
	// confidence measure has been determined through external testing.
	expected = float64(0.958544)
	actual = GetConfidence(50, 52, float64(0.25), float64(0.2), float64(0.01), float64(0.02))
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(0.7072412)
	actual = GetConfidence(100, 200, float64(0.12), float64(0.08), float64(0.1066667), float64(0.07396985))
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(0.983257)
	actual = GetConfidence(1447, 1573, float64(0.1181755), float64(0.08010172), float64(0.2315297), float64(0.1462514))
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestGetDegreesOfFreedom(t *testing.T) {
	// We should get 0 if either count is less than 2.
	expected := float64(0)
	actual := GetDegreesOfFreedom(0, 21, float64(0.25), float64(0.167))
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(0)
	actual = GetDegreesOfFreedom(2, 1, float64(0.25), float64(0.167))
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	// The following tests try various combinations of the parameters whose correct
	// degrees of freedom has been determined through external testing.
	expected = float64(3.84757)
	actual = GetDegreesOfFreedom(3, 3, float64(0.25), float64(0.167))
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(14.95541)
	actual = GetDegreesOfFreedom(10, 21, float64(0.25), float64(0.167))
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(99)
	actual = GetDegreesOfFreedom(100, 30, float64(1.5), float64(0))
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestCalculateTScore(t *testing.T) {
	// We should get a 0 score if the means are the same, regardless of the other
	// parameters.
	expected := float64(0)
	actual := CalculateTScore(100, 101, 2, 2, 1.5, 1.5)
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(0)
	actual = CalculateTScore(500, 600, 3, 3, 0.25, 0.25)
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	// The following tests try various combinations of the parameters whose correct
	// t-score has been determined through external testing.
	expected = float64(-5.36041)
	actual = CalculateTScore(100, 101, 2, 3, 1.5, 2)
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
	expected = float64(-0.07326928)
	actual = CalculateTScore(1000, 1010, 1.999, 2.00001, 0.025, 0.16667)
	if math.Abs(expected-actual) > .0001 {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
