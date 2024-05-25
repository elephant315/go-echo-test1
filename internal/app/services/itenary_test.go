package services

import (
	"reflect"
	"testing"
)

// TestReconstructItinerary tests the ReconstructItinerary function with various inputs
func TestReconstructItinerary(t *testing.T) {
	tests := []struct {
		name        string
		tickets     [][]string
		wantItin    []string
		expectError bool
	}{
		{
			name:        "Normal flow",
			tickets:     [][]string{{"LAX", "DXB"}, {"JFK", "LAX"}, {"SFO", "SJC"}, {"DXB", "SFO"}},
			wantItin:    []string{"JFK", "LAX", "DXB", "SFO", "SJC"},
			expectError: false,
		},
		{
			name:        "Empty input",
			tickets:     [][]string{},
			wantItin:    nil,
			expectError: true,
		},
		{
			name:        "No valid start",
			tickets:     [][]string{{"LAX", "SFO"}, {"SFO", "LAX"}},
			wantItin:    nil,
			expectError: true,
		},
		{
			name:        "Single ticket",
			tickets:     [][]string{{"ATL", "ORD"}},
			wantItin:    []string{"ATL", "ORD"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItin, err := ReconstructItinerary(tt.tickets)
			if (err != nil) != tt.expectError {
				t.Errorf("ReconstructItinerary() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !reflect.DeepEqual(gotItin, tt.wantItin) {
				t.Errorf("ReconstructItinerary() = %v, want %v", gotItin, tt.wantItin)
			}
		})
	}
}
