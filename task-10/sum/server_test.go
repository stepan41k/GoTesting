package task10

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestReadFile(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
	}

	e := httpexpect.Default(t, u.String())

	cases := []struct{
		name string
		x int
		y int
		status int
		expectedErr string
	}{
		{
			name: "Success",
			x: 1,
			y: 1,
			status: http.StatusOK,
		},
		{
			name: "Invalid X",
			x: -1,
			y: 1,
			status: http.StatusBadRequest,
			expectedErr: "value must be positive",
		},
		{
			name: "Invalid Y",
			x: 10,
			y: -10,
			status: http.StatusBadRequest,
			expectedErr: "value must be positive",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			resp := e.POST("/positive-sum").
				WithJSON(Request{
					X: tt.x,
					Y: tt.y,
				}).Expect().Status(tt.status).JSON().Object()

			if tt.expectedErr != "" {
				resp.Value("error").String().IsEqual(tt.expectedErr)

				return
			}
		})
	}
}