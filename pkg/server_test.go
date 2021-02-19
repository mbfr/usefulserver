package usefulserver

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getPage(t *testing.T) {
	test := []struct {
		name string
		url  string
	}{
		{
			name: "no name",
			url:  "http://localhost:8080",
		},
		{
			name: "name",
			url:  "http://localhost:8080?name=mrbean",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			serverUrl, err := url.Parse(tt.url)
			require.NoError(t, err)

			fakeReq := &http.Request{
				URL: serverUrl,
			}

			page, err := getPage(fakeReq, "../templates")
			require.NoError(t, err)
			require.NotEmpty(t, page)
		})
	}
}
