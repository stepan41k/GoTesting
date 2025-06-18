package task7

import (
	"testing"
	"time"
)

func TestTimeFunc(t *testing.T) {
	now := func() time.Time {
		return time.Now().UTC()
	}

	cases := []struct{
		name string
		deadline time.Time
		expected bool
	}{
		{
			name: "Expired",
			deadline: time.Now().UTC().Add(-1 * time.Hour),
			expected: true,
		},
		{
			name: "Not expired",
			deadline: time.Now().UTC().Add(1 * time.Hour),
			expected: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result := IsExprired(now, tt.deadline) 
			
			if result != tt.expected {
				t.Errorf("Expected: %v Actual: %v", tt.expected, result)
			}
		})
	}
}