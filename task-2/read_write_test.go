package task2

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	cases := []struct{
		name string
		path string
		expectedErr string
	}{
		{
			name: "Success",
			path: "./data1.txt",
		},
		{
			name: "Failed to find file",
			path: "./data4.txt",
			expectedErr: "file does not exist",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			_, err := ReadFile(tc.path)
			
			if tc.expectedErr != "" {
				if err.Error() != tc.expectedErr {
					t.Errorf("expected error: %s actual error: %s", tc.expectedErr, err.Error())
				}

				return
			}
		})
	}
}