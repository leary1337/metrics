package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_Update(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}
	tests := []struct {
		name       string
		method     string
		requestURI string
		want       want
	}{
		{
			"valid update metric POST test #1",
			http.MethodPost,
			"/update/counter/someMetric/527",
			want{
				code:        http.StatusOK,
				contentType: "text/plain",
			},
		},
		{
			"method not allowed #2",
			http.MethodGet,
			"/update/counter/someMetric/527",
			want{
				code: http.StatusMethodNotAllowed,
			},
		},
		{
			"metric not found #3",
			http.MethodPost,
			"/update/counter/527",
			want{
				code: http.StatusNotFound,
			},
		},
		{
			"invalid metric value #4",
			http.MethodPost,
			"/update/counter/someMetric/abc",
			want{
				code: http.StatusBadRequest,
			},
		},
		{
			"invalid metric type #5",
			http.MethodPost,
			"/update/someType/someMetric/527",
			want{
				code: http.StatusBadRequest,
			},
		},
	}

	h := &Handler{
		storage: NewMemStorage(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.requestURI, nil)
			w := httptest.NewRecorder()
			h.Update(w, r)

			result := w.Result()
			assert.Equal(t, tt.want.code, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
		})
	}
}
