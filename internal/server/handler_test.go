package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	ts := httptest.NewServer(MetricRouter(&Handler{
		storage: NewMemStorage(),
	}))
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, _ := testRequest(t, ts, tt.method, tt.requestURI, nil)

			require.Equal(t, tt.want.code, response.StatusCode)
			if tt.want.contentType != "" {
				assert.Equal(t, tt.want.contentType, response.Header.Get("Content-Type"))
			}
		})
	}
}

func TestHandler_GetMetricValue(t *testing.T) {
	type want struct {
		code        int
		contentType string
		body        string
	}
	tests := []struct {
		name       string
		method     string
		requestURI string
		want       want
	}{
		{
			"valid get metric test #1",
			http.MethodGet,
			"/value/gauge/testMetric",
			want{
				code:        http.StatusOK,
				contentType: "text/plain",
				body:        "2.32",
			},
		},
		{
			"invalid URL params test #2",
			http.MethodGet,
			"/value/testMetric",
			want{
				code: http.StatusNotFound,
			},
		},
		{
			"unknown metric test #2",
			http.MethodGet,
			"/value/gauge/unknownMetric",
			want{
				code: http.StatusNotFound,
			},
		},
	}

	ts := httptest.NewServer(MetricRouter(&Handler{
		storage: &MemStorage{
			metrics: map[string]Metric{
				"gauge_testMetric": *NewMetric("gauge", "testMetric", 2.32),
			},
		},
	}))
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, body := testRequest(t, ts, tt.method, tt.requestURI, nil)

			require.Equal(t, tt.want.code, response.StatusCode)
			if tt.want.contentType != "" {
				assert.Equal(t, tt.want.contentType, response.Header.Get("Content-Type"))
			}
			if tt.want.body != "" {
				assert.Equal(t, tt.want.body, body)
			}
		})
	}
}

func TestHandler_GetAllMetrics(t *testing.T) {
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
			"get all metrics test #1",
			http.MethodGet,
			"/",
			want{
				code:        http.StatusOK,
				contentType: "text/html",
			},
		},
	}

	h, err := NewHandler(&MemStorage{
		metrics: map[string]Metric{
			"gauge_testMetric": *NewMetric("gauge", "testMetric", 2.32),
		},
	})
	if err != nil {
		return
	}

	ts := httptest.NewServer(MetricRouter(h))
	defer ts.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, body := testRequest(t, ts, tt.method, tt.requestURI, nil)
			require.Equal(t, tt.want.code, response.StatusCode)
			assert.NotEmpty(t, body)
		})
	}
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	require.NoError(t, err)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}
