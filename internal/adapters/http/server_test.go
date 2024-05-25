package http

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHandleItinerary(t *testing.T) {
	e := echo.New()
	tests := []struct {
		name         string
		body         string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "valid input",
			body:         `[["JFK", "LAX"], ["LAX", "DXB"], ["DXB", "SFO"], ["SFO", "SJC"]]`,
			expectedCode: http.StatusOK,
			expectedBody: `["JFK","LAX","DXB","SFO","SJC"]`,
		},
		{
			name:         "invalid input",
			body:         `not a json`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"Invalid input"}`,
		},
		{
			name:         "empty input",
			body:         `[]`,
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"Failed to construct itinerary"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/itinerary", bytes.NewBufferString(tt.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, handleItinerary(c)) {
				assert.Equal(t, tt.expectedCode, rec.Code)
				assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			}
		})
	}
}
