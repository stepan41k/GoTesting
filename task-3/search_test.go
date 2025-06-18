package task3

import "testing"

func TestReadFile(t *testing.T) {
	cases := []struct{
		name string
		arr []int
		target int
		expectedErr string
	}{
		{
			name: "Success",
			arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target: 3,
		},
		{
			name: "Failed to find value",
			arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target: 11,
			expectedErr: ErrCantFindTarget.Error(),
		},
		{
			name: "Empty array",
			arr: []int{},
			target: 10,
			expectedErr: ErrEmptyArr.Error(),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			_, err := BinSearch(tc.arr, tc.target)
			
			if tc.expectedErr != "" {
				if err.Error() != tc.expectedErr {
					t.Errorf("expected error: %s actual error: %s", tc.expectedErr, err.Error())
				}

				return
			}
		})
	}
}