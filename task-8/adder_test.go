package task8

import (
	"sync"
	"testing"
)

func TestInsertRead(t *testing.T) {
	cases := []struct {
		name string
		values      []int
		expectedErr string
	}{
		{
			name: "Success",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "Empty",
			values: []int{},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			mm := New()
			wg := &sync.WaitGroup{}

			for i := 0; i < len(tt.values); i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					mm.InsertValue(tt.values[i], tt.values[i] * 10)
				}()
			}

			wg.Wait()

			for i := 0; i < len(tt.values); i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					mm.ReadValue(tt.values[i])
				}()
			}

			wg.Wait()
		})
	}
}
