package display

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendName(t *testing.T) {
	cases := []struct {
		name        string
		firstName   string
		statusCode  string
		expectedErr string
	}{
		{
			name:      "Success",
			firstName: "Arnold",
		},
		{
			name:        "Failed to display the user",
			firstName:   "",
			expectedErr: "empty name",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/user?name=%s", tc.firstName)
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(tc.firstName)))
			require.NoError(t, err)

			rr := httptest.NewRecorder()

			DisplayName(rr, req)
			body := rr.Body.String()

			var resp Response

			require.NoError(t, json.Unmarshal([]byte(body), &resp))
			


			require.Equal(t, tc.expectedErr, resp.Error)
		})
	}
}
