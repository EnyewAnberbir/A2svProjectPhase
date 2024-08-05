package main

import (
	"testing"
)
func TestAverageGrades(t *testing.T) {
	tests := []struct {
		name          string
		subjects      map[string]float64
		expectedAvg   float64
		expectedError bool
	}{
		{
			name:          "No subjects",
			subjects:      map[string]float64{},
			expectedAvg:   0,
			expectedError: true,
		},
		{
			name: "Single subject",
			subjects: map[string]float64{
				"Applied": 90,
			},
			expectedAvg:   90,
			expectedError: false,
		},
		{
			name: "Multiple subjects",
			subjects: map[string]float64{
				"Math":    90,
				"English": 80,
				"Physics": 85,
			},
			expectedAvg:   85,
			expectedError: false,
		},
		{
			name: "Invalid grade values",
			subjects: map[string]float64{
				"Math":    -10,
				"English": 200,
			},
			expectedAvg:   95,
			expectedError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			avg, err := averageGrades(test.subjects)
			if (err != nil) != test.expectedError {
				t.Errorf("expected error: %v, got: %v", test.expectedError, err)
			}
			if avg != test.expectedAvg {
				t.Errorf("expected average: %.2f, got: %.2f", test.expectedAvg, avg)
			}
		})
	}
}
