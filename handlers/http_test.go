package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestHTTPHandler_Add(t *testing.T) {
	assertRoute(t, "GET", "/boink?a=1&b=3", http.StatusNotFound, "404 page not found\n")
	assertRoute(t, "POST", "/add?a=1&b=3", http.StatusMethodNotAllowed, "Method Not Allowed\n")
	assertRoute(t, "GET", "/add?a=NaN&b=3", http.StatusUnprocessableEntity, "The a parameter must be an integer\n")
	assertRoute(t, "GET", "/add?a=1&b=NaN", http.StatusUnprocessableEntity, "The b parameter must be an integer\n")
	assertRoute(t, "GET", "/add?a=1&b=3", http.StatusOK, "4")
	assertRoute(t, "GET", "/sub?a=3&b=1", http.StatusOK, "2")
	assertRoute(t, "GET", "/mul?a=3&b=2", http.StatusOK, "6")
	assertRoute(t, "GET", "/div?a=30&b=2", http.StatusOK, "15")
}

func assertRoute(t *testing.T, method, target string, statusCode int, responseBody string) {
	t.Run(fmt.Sprintf("%s %s", method, target), func(t *testing.T) {
		request := httptest.NewRequest(method, target, nil)
		recorder := httptest.NewRecorder()
		var logBuffer bytes.Buffer
		router := NewHTTPRouter(&logBuffer)

		requestDump, err := httputil.DumpRequest(request, true)
		assertErr(t, err, nil)
		t.Logf("request dump:\n%s", string(requestDump))

		router.ServeHTTP(recorder, request)

		responseDump, err := httputil.DumpResponse(recorder.Result(), true)
		assertErr(t, err, nil)
		t.Logf("response dump:\n%s", string(responseDump))

		t.Log(logBuffer.String())

		assertEqual(t, recorder.Code, statusCode)
		assertEqual(t, recorder.Body.String(), responseBody)
	})
}
