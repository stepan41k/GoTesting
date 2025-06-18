package task1

import "testing"

func TestSumOfTwo(t *testing.T) {
	cases := []struct{
		name string
		x int
		y int
		expectedValue int
	}{
		{
			name: "test1",
			x: 5,
			y: 5,
			expectedValue: 10,
		},
		{
			name: "test2",
			x: 20,
			y: 0,
			expectedValue: 20,
		},
		{
			name: "test3",
			x: 15,
			y: 15,
			expectedValue: 30,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			res := SumOfTwo(tt.x, tt.y)
			if res != tt.expectedValue{
				t.Errorf("Expected value: %d Actual value: %d", tt.expectedValue, res)
			}
		})
	}
}