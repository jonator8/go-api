package controllers

import (
	"github.com/jonator8/go-api/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	app := internal.NewApp(nil, &internal.CommandsBootstrap{})

	testCases := []struct {
		desc string
		body string
		want struct {
			code     int
			response string
		}
	}{
		{
			desc: "fails because empty fields",
			body: `{"title":"", "body":""}`,
			want: struct {
				code     int
				response string
			}{
				code:     400,
				response: `{"code":400,"err":"field 'Title' with value '' cannot be empty, field 'Body' with value '' cannot be empty"}`,
			},
		},
		{
			desc: "fails because malformed body",
			body: `{title:"", "body":""}`,
			want: struct {
				code     int
				response string
			}{
				code:     400,
				response: `{"code":400,"err":"create_new_cotroller: error formating json: invalid character 't' looking for beginning of object key string"}`,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.body)
			r := httptest.NewRequest(http.MethodPost, "/api/v1/news", reader)
			w := httptest.NewRecorder()

			CreateNewController(app).ServeHTTP(w, r)
			body, err := io.ReadAll(w.Result().Body)

			assert.Nil(t, err, "should be nil")
			assert.Equal(t, tc.want.code, w.Code)
			assert.Equal(t, tc.want.response, string(body))
		})
	}
}
